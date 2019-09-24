package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Schema struct {
	PostgresDB struct {
		Username string
		Password string
		Host     string
		Port     int
		Debug    bool
	}

	Paging struct {
		Limit string
	}

	Mongo struct {
		URI      string `mapstructure:"uri"`
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mongo"`

	Encryption struct {
		EncryptionKey string
		EncryptSecret string
	}
}

var Config Schema

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&Config)
}
