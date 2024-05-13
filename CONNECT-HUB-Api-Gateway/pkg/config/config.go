package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Port                string `mapstructure:"PORT"`
	ConnetHubAuth       string `mapstructure:"Connect_Hub_Auth"`
	ConnetHubPost       string `mapstructure:"Connect_Hub_Post"`
	ConnetHubChat       string `mapstructure:"Connect_Hub_Chat"`
	ConnetHubJob        string `mapstructure:"Connect_Hub_Job"`
	JobSeekerAccessKey  string `mapstructure:"Job_Seeker_AccessKey"`
	AdminAccessKey      string `mapstructure:"Admin_AccessKey"`
	RecruiterAccessKey  string `mapstructure:"Recruiter_AccessKey"`
	JobSeekerRefreshKey string `mapstructure:"Job_Seeker_RefreshKey"`
	AdminRefreshKey     string `mapstructure:"Admin_RefreshKey"`
	RecruiterRefreshKey string `mapstructure:"Recruiter_RefreshKey"`
}

var envs = []string{
	"PORT", "Connect_Hub_Auth", "Connect_Hub_Post",
	"Connect_Hub_Chat", "Job_Seeker_AccessKey",
	"Admin_AccessKey", "Recruiter_AccessKey",
	"Job_Seeker_RefreshKey", "Admin_RefreshKey",
	"Recruiter_RefreshKey", "Connect_Hub_Job",
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
