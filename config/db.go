package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     uint   `env:"DB_PORT"`
	SSLMode  string `env:"DB_SSL_MODE"`
}

func NewDBConfig(v *viper.Viper) (*DBConfig, error) {
	c := new(DBConfig)
	if err := v.Unmarshal(c, envDecoder); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *DBConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.DBName)
}
