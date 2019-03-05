package main

import (
	"github.com/j7mbo/goenvconfig"
	"fmt"
)

type Config struct {
	host     string  `env:"HOME" default:"localhost"`
	port     int     `env:"PORT" default:"8080"`
}

func (c *Config) GetHost() string { return c.host }
func (c *Config) GetPort() int { return c.port }

func main() {
	config := Config{}
	parser := goenvconfig.NewGoEnvParser()

	if err := parser.Parse(&config) {
		panic(err)
	}

	fmt.Println(config.GetHost()) // localhost
	fmt.Println(config.GetPort()) // 1337
}
