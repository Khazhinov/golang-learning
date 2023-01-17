package main

import (
	"golang-learning/config"
	"golang-learning/internal/foundation/configurator"
	"golang-learning/internal/helper"
)

func main() {
	config := config.NewConfig()
	configurator.ReadYaml(config)
	configurator.ReadEnvironment(config)

	helper.DD(config.Database.Host)
}
