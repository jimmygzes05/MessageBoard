package config

import (
	"fmt"
	"log"
	"os"

	"github.com/naoina/toml"
)

var Conf *Config

func Load() *Config {
	root, _ := os.Getwd()

	defaultPath := root + "/config/default.toml"
	defaultToml, readErr := os.ReadFile(defaultPath)
	if readErr != nil {
		log.Printf("[ Config Load readErr ] - Err : [ %s ]", readErr)
		return nil
	}

	unmarErr := toml.Unmarshal(defaultToml, &Conf)
	if unmarErr != nil {
		panic(fmt.Sprintf("[ 🦁 Config Load Unmarshal ] - Err : [ %s ]", unmarErr))
	}

	return Conf
}
