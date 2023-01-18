package configurator

import (
	"fmt"
	"github.com/fatih/structs"
	"golang-learning/config"
	"golang-learning/internal/foundation/cmd"
	"golang-learning/internal/foundation/filesystem"
	"golang-learning/internal/helper"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

func ReadDefaultValues(configFile Configurable) {
	err := yaml.Unmarshal(config.DefaultValues, configFile)
	cmd.CheckError(err)
}

func GetDefaultValues() []byte {
	defaultValuesDir := filesystem.BasePath(fmt.Sprintf("%s%c%s", "config", os.PathSeparator, "values"))
	descriptor, err := os.Open(defaultValuesDir)
	if err != nil {
		cmd.CheckError(err)
	}
	files, err := descriptor.Readdir(0)
	if err != nil {
		cmd.CheckError(err)
	}

	base := map[string]interface{}{}

	for _, file := range files {
		filepath := fmt.Sprintf("%s%c%s", defaultValuesDir, os.PathSeparator, file.Name())

		currentMap := map[string]interface{}{}
		file, err := os.Open(filepath)
		if err != nil {
			cmd.CheckError(err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				cmd.CheckError(err)
			}
		}(file)

		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(currentMap)
		if err != nil {
			cmd.CheckError(err)
		}

		// merge both yaml data recursively
		base = helper.MergeMaps(base, currentMap)
	}

	result, err := yaml.Marshal(base)
	cmd.CheckError(err)

	return result
}

func ReadYaml(config Configurable) {
	files := make([]string, 0)

	for _, name := range structs.Names(config) {
		filename := fmt.Sprintf("%s.yaml", strings.ToLower(name))
		files = append(files, filename)
	}

	for _, filename := range files {
		filepath := filesystem.BasePath(fmt.Sprintf("%s%c%s%c%s", "config", os.PathSeparator, "values", os.PathSeparator, filename))

		if _, err := os.Stat(filepath); err == nil {
			file, err := os.Open(filepath)
			if err != nil {
				cmd.CheckError(err)
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					cmd.CheckError(err)
				}
			}(file)

			decoder := yaml.NewDecoder(file)
			err = decoder.Decode(config)
			if err != nil {
				cmd.CheckError(err)
			}
		}
	}
}
