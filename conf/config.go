package conf

import (
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() *viper.Viper {
	MODE := os.Getenv("MODE")

	v := viper.New()

	v.AddConfigPath("./conf/")
	v.SetConfigType("yaml")

	if MODE == "PRODUCTION" {
		v.SetConfigName("production-config")
	} else if MODE == "DEVELOPMENT" {
		v.SetConfigName("development-config")
	} else {
		v.SetConfigName("config")
	}

	if err := v.ReadInConfig(); err != nil {
		return nil
	}

	return v
}
