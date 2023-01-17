package config

type Config struct {
	Database    DatabaseConnectionConfig `yaml:"database"`
	Application ApplicationConfig
}

func NewConfig() *Config {
	return &Config{}
}

func (config *Config) Configurable() {}
