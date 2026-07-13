package main

import (
	"caplet/internal/cli"
	"log"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
