package main

import (
	_ "bhi/routers"
	"context"
	"io/ioutil"

	"net/http"

	"github.com/aws/aws-xray-sdk-go/awsplugins/ec2"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/net/context/ctxhttp"
)

func init() {
	ec2.Init()
	xray.Configure(xray.Config{
		DaemonAddr:       "127.0.0.1:2000", // default
		ServiceVersion:   "1.2.3",
	})
}

func middleware(componentName string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return xray.Handler(xray.NewFixedSegmentNamer(componentName), h)
	}
}

func main() {
	logs.Info("In startup.")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	ctx, root := xray.BeginSegment(context.Background(), "get my ip")
	resp, err := ctxhttp.Get(ctx, xray.Client(nil), "https://ipv4.kabakaev.com/")
    if err != nil {
      logs.Warn( err )
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logs.Warn(err)
		}
		bodyString := string(bodyBytes)
		logs.Info("My IP is %s", bodyString)
	}

	root.Close(nil)

	//beego.Run()
	beego.RunWithMiddleWares(":8080", middleware("bhi"))
}

//beego.App.Run >> http.ListenAndServe(addr, reqidMiddleware(dlogMiddleware(mux)))
// beego.RunWithMiddleWares (addr, mws ... MiddleWare)

//func main() {
//	http.Handle("/", xray.Handler(xray.NewFixedSegmentNamer("myApp"), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//	  w.Write([]byte("Hello!"))
//	})))
//
//	http.ListenAndServe(":8000", nil)
//  }
