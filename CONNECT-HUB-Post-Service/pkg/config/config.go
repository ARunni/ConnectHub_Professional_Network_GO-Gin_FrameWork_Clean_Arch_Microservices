package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Port               string `mapstructure:"PORT"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBName             string `mapstructure:"DB_NAME"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	Connect_Hub_Auth   string `mapstructure:"Connect_Hub_Auth"`
	Connect_Hub_Post   string `mapstructure:"Connect_Hub_Post"`
	Connect_Hub_Chat   string `mapstructure:"Connect_Hub_Chat"`
	AWSRegion          string `mapstructure:"AWSRegion"`
	AWSAccesskeyID     string `mapstructure:"AWSAccesskeyID"`
	AWSSecretaccesskey string `mapstructure:"AWSSecretaccesskey"`
	BucketName         string `mapstructure:"BucketName"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER",
	"DB_PORT", "DB_PASSWORD", "PORT",
	"Connect_Hub_Auth", "Connect_Hub_Post",
	"Connect_Hub_Chat", "AWSRegion",
	"AWSAccesskeyID", "AWSSecretaccesskey",
	"BucketName",
}

func LoadConfig() (Config, error) {
	var config Config
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil

}
