package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Schema struct {
	PostgresDB struct {
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		DatabaseName     string `mapstructure:"db_name"`
		Password string `mapstructure:"password"`
		Debug    bool   `mapstructure:"debug"`
	}

	Paging struct {
		Limit int
	}

	Mongo struct {
		URI      string `mapstructure:"uri"`
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DatabaseName string `mapstructure:"db_name"`
	} `mapstructure:"mongo"`

	Encryption struct {
		JWTSecret        string `mapstructure:"jwt_secret"`
		JWTSecretRefresh string `mapstructure:"jwt_secret_refresh"`
		JWTExp           int    `mapstructure:"jwt_exp"`
		JWTExpRefresh           int    `mapstructure:"jwt_exp_refresh"`
	} `mapstructure:"encryption"`
}

var Config Schema

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")
	viper.AddConfigPath("../config")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&Config)
}
