package config

import "github.com/spf13/viper"

type Config struct {
	ServerLocalHost string `mapstructure:"ServerLocalhost"`
	DBURL           string `mapstructure:"DBURL"`
}

func LoadConfig(path string, file string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
