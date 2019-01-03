package main

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Environment struct {
	Home string `env:"HOME"`

	Jenkins struct {
		BuildId     *string `env:"BUILD_ID"`
		BuildNumber int     `env:"BUILD_NUMBER"`
		Ci          bool    `env:"CI"`
	}

	Extras env.EnvSet
}

func main() {
	var environment Environment
	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}
	// Remaining environment variables.
	environment.Extras = es

	// ...

	es, err = env.Marshal(environment)
	if err != nil {
		log.Fatal(err)
	}

	home := "/tmp/edgarl"
	cs := env.ChangeSet{
		"HOME":         &home,
		"BUILD_ID":     nil,
		"BUILD_NUMBER": nil,
	}
	es.Apply(cs)

	environment = Environment{}
	err = env.Unmarshal(es, &environment)
	if err != nil {
		log.Fatal(err)
	}

	environment.Extras = es
}
