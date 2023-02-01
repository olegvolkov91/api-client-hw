package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

const pk string = "PRIMARY_TOKEN"

type Config struct {
	PrimaryToken string
	ClientAddr   string `mapstructure:"client_addr"`

	Messages
}

type Messages struct {
	Errors
}

type Errors struct {
	UnableToLoad       string `mapstructure:"unable_to_load"`
	UnableToParse      string `mapstructure:"unable_to_parse"`
	UnableToConvert    string `mapstructure:"unable_to_convert"`
	SomethingWentWrong string `mapstructure:"something_went_wrong"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := godotenv.Load(); err != nil {
		return nil, errors.New("no .env file found")
	}
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors); err != nil {
		return nil, err
	}

	if err := getEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getEnv(cfg *Config) error {
	token, exists := os.LookupEnv(pk)

	if !exists {
		return errors.New("primary token does not exist in .env file")
	}
	cfg.PrimaryToken = token
	return nil
}
