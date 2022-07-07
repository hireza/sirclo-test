package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	DBName string
	DBUser string
	DBPass string
}

func NewConfig(path string) (*Config, error) {
	viper.AddConfigPath("packages/config")
	if path != "" {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("error when read config")
		return nil, err
	}

	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Error().Err(err).Msg("error when unmarshal config")
		return nil, err
	}

	return config, nil
}
