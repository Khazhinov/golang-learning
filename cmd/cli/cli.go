package main

import (
	"fmt"
	"golang-learning/config"
	"golang-learning/internal/foundation/configurator"
	"os"
)

func main() {
	config := config.NewConfig()
	configurator.ReadDefaultValues(config)
	configurator.ReadEnvironment(config)

	fmt.Println(config.Database.Host)

	os.Exit(0)
}
