package controllers

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"

	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

type StatGraphController struct {
	beego.Controller
}

func (c *StatGraphController) Get() {
	c.EnableRender = false

	c.Ctx.Output.Header("Content-Type", "image/png")

	rb, err := renderStatPNG()
	if err != nil {
		logs.Error("Cannot render PNG: %w", err)
	}
	if err := c.Ctx.Output.Body(rb); err != nil {
		logs.Error("Cannot set body: %w", err)
	}
}

func renderStatPNG() ([]byte, error) {
	goodStyle := chart.Style{
		FillColor:   drawing.ColorFromHex("13c158"),
		StrokeColor: drawing.ColorFromHex("13c158"),
		StrokeWidth: 0,
	}

	badStyle := chart.Style{
		FillColor:   drawing.ColorFromHex("c11313"),
		StrokeColor: drawing.ColorFromHex("c11313"),
		StrokeWidth: 0,
	}

	bars := []chart.Value{}
	for i := 32; i <= 128; i = i+16 { // From 16 to 64 in steps of 8.
		var val float64
		intVal, ok := responses.Get(fmt.Sprint(i)).(int)
		if ok {
			val = float64(intVal)
		} else { // Prevent zero Y-scale.
			val = rand.Float64() / 100
		}

		var style chart.Style
		if val >= 0 {
			style = goodStyle
		} else {
			style = badStyle
		}
		bars = append(bars, chart.Value{Value: val, Style: style, Label: fmt.Sprint(i)})
		logs.Info("added $i = $d", i, val)
	}

	sbc := chart.BarChart{
		Title: "Preferred header size",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		//YAxis: chart.YAxis{
		//	Ticks: []chart.Tick{
		//		{Value: -4.0, Label: "-4"},
		//		{Value: -2.0, Label: "-2"},
		//		{Value: 0, Label: "0"},
		//		{Value: 2.0, Label: "2"},
		//		{Value: 4.0, Label: "4"},
		//		{Value: 6.0, Label: "6"},
		//		{Value: 8.0, Label: "8"},
		//		{Value: 10.0, Label: "10"},
		//		{Value: 12.0, Label: "12"},
		//	},
		//},
		UseBaseValue: true,
		BaseValue:    0.0,
		Bars:         bars,
	}

	var buf bytes.Buffer
	err := sbc.Render(chart.PNG, &buf)
	return buf.Bytes(), err
}
