package config

// Server is the configuration for the server
type Server struct {
	Port int `mapstructure:"port" env:"SERVER_PORT"`
}
