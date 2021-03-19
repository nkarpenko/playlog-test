package conf

import (
	"github.com/spf13/viper"
)

// Config global var.
var Config *Configuration

// Configuration of the service
type Configuration struct {
	Name         string   `mapstructure:"name"`
	Desc         string   `mapstructure:"desc"`
	Host         string   `mapstructure:"host"`
	Port         int      `mapstructure:"port"`
	Version      string   `mapstructure:"version"`
	ShutdownTime int      `mapstructure:"sht"`
	Database     DBConfig `mapstructure:"db"`
}

// Load the config file
func Load(file string) (*Configuration, error) {
	var config *Configuration
	viper.SetConfigType("yaml")
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		return config, err
	}

	config = new(Configuration)
	if err := viper.Unmarshal(config); err != nil {
		return config, err
	}

	return config, nil
}
