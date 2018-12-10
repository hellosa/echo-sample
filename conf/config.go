package conf

import (
	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	v := viper.New()

	v.AddConfigPath("./conf/")

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return nil
	}

	return v
}
