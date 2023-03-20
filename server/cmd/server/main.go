package main

import (
	"flag"
	"fmt"
	"log"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}
