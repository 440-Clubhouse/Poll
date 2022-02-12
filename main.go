package main

import (
    "flag"

    "github.com/440-clubhouse/poll/app"
)

func main() {
    var configPath string
    flag.StringVar(&configPath, "config", "./config.json", "The config path")
    flag.StringVar(&configPath, "c", "./config.json", "The short alias of -config")
    flag.Parse()

    app.Init(configPath)
}
