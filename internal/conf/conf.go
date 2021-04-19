package conf

import "github.com/spf13/viper"

type Config struct {
	ServerLocalHost string `mapstructure:"SERVER_LOCALHOST"`
	DB_URL string `mapstructure:"DB_URL"`

}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
