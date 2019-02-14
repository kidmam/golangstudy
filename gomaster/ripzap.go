package main

import (
	"github.com/bloom42/rz-go"
	"os"

	"github.com/bloom42/rz-go/log"
)

func main() {

	env := os.Getenv("GO_ENV")
	hostname, _ := os.Hostname()

	// update global logger's context fields
	log.Logger = log.Config(rz.With(func(e *rz.Event) {
		e.String("hostname", hostname).
			String("environment", env)
	}))

	if env == "production" {
		log.Logger = log.Config(rz.Level(rz.InfoLevel))
	}

	log.Info("info from logger", func(e *rz.Event) {
		e.String("hello", "world")
	})
	// {"level":"info","hostname":"","environment":"","hello":"world","timestamp":"2019-02-07T09:30:07Z","message":"info from logger"}
}
