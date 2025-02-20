package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	DbPath string
}

func LoadConfiguration(configPath, configName, configType string) (*Configuration, error) {
	var config *Configuration

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("LoadConfiguration failed: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("LoadConfiguration failed: %w", err)
	}

	return config, nil
}