package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

const (
	slackEventTimeout = 3 * time.Second
)

var (
	flagConfigPath = flag.String("dan-demand.config", "", "Configuration file location")
)

func main() {
	flag.Parse()

	config, err := LoadConfig(*flagConfigPath)
	if err != nil {
		glog.Fatal(errors.Wrap(err, "failed to load config: "))
	}
	glog.Info("Loaded config: ", config)

	engine, err := NewEngine(config)
	if err != nil {
		glog.Fatal(errors.Wrap(err, "failed to create Engine: "))
	}

	glog.Infof("DanDemand running on %s", config.Server.Address)
	glog.Fatal(engine.ListenAndServe())
}
