package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// config holds config variables
type config struct {
	databaseAddress string // connection address of the database
}

// DatabaseAddress returns database address
func (c config) DatabaseAddress() string {
	return c.databaseAddress
}

// GetConfig reads content of the input config file and puts it into a config instance and returns it
func GetConfig(configName string, configPath string, configType string) *config {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.AddConfigPath(".")

	// reading environment variables
	viper.AutomaticEnv()

	// read configs
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("unable to read config file: %w", err))
	}

	configs := config{
		databaseAddress: viper.Get("DATABASE_ADDRESS").(string),
	}

	return &configs
}
