package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerHost string `mapstructure:"serverHost"`
	DBURL      string `mapstructure:"dbUrl"`
	Front      string `mapstructure:"front"`
	Mail       MailConfig
}

type MailConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Sender       string `mapstructure:"sender"`
	AuthEmail    string `mapstructure:"user"`
	AuthPassword string `mapstructure:"password"`
}

func LoadConfig(file string) (config Config, err error) {
	viper.SetConfigFile(file)
	viper.AutomaticEnv()
	err = viper.BindEnv("mail.sender", "MAIL_SENDER")
	if err != nil {
		return
	}

	err = viper.BindEnv("mail.user", "MAIL_USER")
	if err != nil {
		return
	}

	err = viper.BindEnv("mail.password", "MAIL_PASSWORD")
	if err != nil {
		return
	}

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
