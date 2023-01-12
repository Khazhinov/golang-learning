package main

import (
	"golang-learning/configs"
	"golang-learning/internal/foundation/configurator"
	"golang-learning/internal/helper"
)

func main() {
	config := configs.NewConfig()
	configurator.ReadYaml(config)
	configurator.ReadEnvironment(config)

	helper.DD(config.Database.Host)
}
