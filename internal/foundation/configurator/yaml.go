package configurator

import (
	"fmt"
	"github.com/fatih/structs"
	"golang-learning/internal/foundation/cmd"
	"golang-learning/internal/foundation/filesystem"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

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
