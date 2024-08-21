package config

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	QueueUser         string `mapstructure:"QUEUE_USER"`
	QueuePassword     string `mapstructure:"QUEUE_PASSWORD"`
	QueueHost         string `mapstructure:"QUEUE_HOST"`
	QueuePort         string `mapstructure:"QUEUE_PORT"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfig(path string) (*config, error) {
	viper.SetConfigName("app_config")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	// Enviroment variables will overide the .env values eg export DB_DRIVER=postgres will overide the DB_DRIVER value in .env
	// enabling us to change some values without changing the .env file for testing purposes	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
