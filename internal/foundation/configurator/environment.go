package configurator

import (
	"github.com/kelseyhightower/envconfig"
	"golang-learning/internal/foundation/cmd"
)

func ReadEnvironment(config Configurable) {
	err := envconfig.Process("", config)
	if err != nil {
		cmd.CheckError(err)
	}
}
