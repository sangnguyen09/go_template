package config

import (
	"fmt"

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

	Encryption struct {
		EncryptionKey string
		EncryptSecret string
	}
}

var Config Schema

func init() {
	viper.SetConfigName("config")        // name of config file (without extension)
	viper.AddConfigPath("/etc/appname/") // path to look for the config file in
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&Config)
}
