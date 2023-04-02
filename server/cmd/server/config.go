package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Client      string `json:"client"`
	FilesFolder string `json:"files_folder"`
}

func NewConfig() (Config, error) {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("fatal error config file: %w", err)
	}

	return Config{
		Client:      viper.GetString("client"),
		FilesFolder: viper.GetString("files_folder"),
	}, nil
}
