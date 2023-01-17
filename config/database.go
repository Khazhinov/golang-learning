package config

type DatabaseConnectionConfig struct {
	Host     string `yaml:"host" envconfig:"DB_HOST"`
	Port     string `yaml:"port" envconfig:"DB_PORT"`
	Username string `yaml:"username" envconfig:"DB_USERNAME"`
	Password string `yaml:"password" envconfig:"DB_PASSWORD"`
	Database string `yaml:"database" envconfig:"DB_DATABASE"`
}
