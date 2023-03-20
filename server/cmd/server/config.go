package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	BackendAPI  string `json:"backend_api"`
	FilesFolder string `json:"files_folder"`
}

func NewConfig() (Config, error) {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("fatal error config file: %w", err)
	}

	return Config{
		BackendAPI:  viper.GetString("backend_api"),
		FilesFolder: viper.GetString("files_folder"),
	}, nil
}
