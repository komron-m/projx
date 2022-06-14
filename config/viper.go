package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func envDecoder(c *mapstructure.DecoderConfig) {
	c.TagName = "env"
}

func NewViper() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(".env") // name of config file (without extension)
	v.AddConfigPath(".")    // optionally look for config in the working directory

	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		return nil, err
	}
	return v, nil
}
