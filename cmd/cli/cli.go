package main

import (
	"fmt"
	"golang-learning/config"
	"golang-learning/internal/foundation/configurator"
	"os"
)

func main() {
	config := config.NewConfig()
	//configurator.ReadYaml(config)
	configurator.ReadDefaultValues(config)

	fmt.Println(config.Database.Host)

	os.Exit(0)
}
