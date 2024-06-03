package config

import "github.com/spf13/viper"

type Config struct {
	DBUri            string `mapstructure:"DB_URI"`
	Port             string `mapstructure:"PORT"`
	Connect_Hub_Auth string `mapstructure:"Connect_Hub_Auth"`
	Connect_Hub_Post string `mapstructure:"Connect_Hub_Post"`
	Connect_Hub_Job  string `mapstructure:"Connect_Hub_Job"`
	KafkaBrokers     string `mapstructure:"KAFKA_BROKERS"`
	KafkaTopic       string `mapstructure:"KAFKA_TOPIC"`
}

var envs = []string{
	"DB_URI", "PORT", "KAFKA_BROKERS", "KAFKA_TOPIC",
	"Connect_Hub_Auth", "Connect_Hub_Post",
	"Connect_Hub_Job",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
