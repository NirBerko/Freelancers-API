package app

import (
	"github.com/spf13/viper"
	"log"
)

var Config *appConfig

type appConfig struct {
	DSN    string
	Mode   string
	Server struct {
		Port int
	}
	Jwt struct {
		JWTSigningKey      string `mapstructure:"signing_key"`
		JWTVerificationKey string `mapstructure:"verification_key"`
	} `mapstructure:"jwt"`
}

func LoadConfig() *appConfig {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetDefault("Server.Port", 8080)

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := v.Unmarshal(&Config)

	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return Config
}
