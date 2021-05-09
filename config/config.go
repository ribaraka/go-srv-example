package config

import "github.com/spf13/viper"

type Config struct {
	ServerLocalHost string `mapstructure:"ServerLocalhost"`
	DBURL           string `mapstructure:"DBURL"`
	StaticAssets    string `mapstructure:"StaticAssets"`
}

func LoadConfig(file string) (config Config, err error) {
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
