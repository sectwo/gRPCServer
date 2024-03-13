package main

import (
	"flag"
	"grpc-server/cmd"
	"grpc-server/config"
)

var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)
	cmd.NewApp(cfg)
}
