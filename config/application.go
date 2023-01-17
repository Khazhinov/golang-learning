package config

type ApplicationConfig struct {
	Host string `envconfig:"APP_HOST"`
	Port string `envconfig:"APP_PORT"`
}
