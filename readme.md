# Beego test app

[Beego](https://beego.me/) is an MVC framework written in Go.

App name `bhi` stands for `Beego Hello`.

## What's inside

The code is generated by `bee new bhi` command.

The index page prints "Hello $NAME", with $NAME coming from environment variable `NAME`.

The app also collects user feedback data and generates a PNG chart out of it.

## Run locally

bee run

## Deploy to production

Upload the binary:
```bash
bee pack && rsync -avH ./bhi.tar.gz  remote.host.name:
```

Setup host:
```bash
cat <<'EOF' | sudo tee /etc/systemd/system/beego.service
# see https://beego.me/docs/deploy/systemctl.md

[Unit]
Description=beego
AssertPathExists=/home/some-user/bhi/bhi

[Service]
Environment="NAME=Username"
WorkingDirectory=/home/some-user/bhi
ExecStart=/home/some-user/bhi/bhi

ExecReload=/bin/kill -HUP $MAINPID
LimitNOFILE=65536
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable beego

tar xvf $HOME/bhi.tar.gz -C $HOME/bhi/ --exclude=conf/app.conf
cat <<EOF > $HOME/bhi/conf/app.conf
appname = bhi
httpport = 8080
runmode = prod
EOF
sudo systemctl start beego
```
