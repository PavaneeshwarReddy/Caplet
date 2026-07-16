package main

import (
	"caplet/internal/cli"
	"caplet/internal/config"
	"log"
)

func main() {

	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
