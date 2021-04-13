package conf

import "github.com/spf13/viper"

type Config struct {
	ServerHost string `mapstructure:"SERVER_HOST"`
	DBSource string `mapstructure:"DB_SOURCE"`
	//DBHost     string `mapstructure:"DB_HOST"`
	//DBPort     string `mapstructure:"DB_PORT"`
	//DBUserName string `mapstructure:"DB_USERNAME"`
	//DBPassword string `mapstructure:"DB_PASSWORD"`
	//DBName     string `mapstructure:"DB_NAME"`

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
