package configurator

import (
	"github.com/kelseyhightower/envconfig"
	"golang-learning/configs"
	"golang-learning/internal/foundation/cmd"
)

func ReadEnvironment(config *configs.Config) {
	err := envconfig.Process("", config)
	if err != nil {
		cmd.CheckError(err)
	}
}
