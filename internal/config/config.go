package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config is the configuration for the whole application
type Config struct {
	Env    string `mapstructure:"env"`
	Server Server `mapstructure:"server"`
}

func LoadConfig(filePath string) (Config, error) {
	var cfg Config

	viper.SetEnvPrefix("githeim")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("env", "prod")
	viper.SetDefault("server.port", 8080)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filePath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
		} else {
			return cfg, err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
