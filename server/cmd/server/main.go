package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"

	"minisearch/server/packages/api"
	"minisearch/server/packages/searchalg"
	"minisearch/server/packages/watcher"
)

var (
	TxtTypes   = regexp.MustCompile(".(txt)$")
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "/etc/app/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)

	watcherInstance := watcher.New(watcher.Settings{
		Folder:    config.FilesFolder,
		FileTypes: TxtTypes,
	})

	err = watcherInstance.RunFilesChecking()
	if err != nil {
		log.Fatal(err)
	}

	err = watcherInstance.CacheData()
	if err != nil {
		log.Fatal(err)
	}

	searchInstance := searchalg.New(searchalg.Settings{}, watcherInstance)
	apiInstance := api.New(api.Settings{}, searchInstance)

	err = apiInstance.ServiceStart()
	if err != nil {
		log.Fatal(err)
	}
}
