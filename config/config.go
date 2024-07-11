package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type AppConfiguration struct {
	Name    string `mapstructure:"APP_NAME"`
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort int    `mapstructure:"APP_PORT"`
}

type DatabaseConfig struct {
	LearningName string `mapstructure:"DB_LEARNING"`
	Port         int    `mapstructure:"DB_PORT"`
	Host         string `mapstructure:"DB_HOST"`
	Username     string `mapstructure:"DB_USER"`
	Password     string `mapstructure:"DB_PASSWORD"`
}

var appConfig *AppConfiguration
var dbConfig *DatabaseConfig
var path string = "./"
var configType string = "env"
var configName string = ".env"

var mu sync.Mutex

func GetDatabaseConfig() *DatabaseConfig {
	mu.Lock()

	defer mu.Unlock()

	if dbConfig == nil {
		viper.AddConfigPath(path)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

		// Init config
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("[Config] - Failed to read config %v", err)
		}

		var dbCfg DatabaseConfig

		if err := viper.Unmarshal(&dbCfg); err != nil {
			panic("unable to unmarshal config")
		}

		dbConfig = &dbCfg
	}

	return dbConfig
}

func GetAppConfiguration() *AppConfiguration {

	mu.Lock()

	defer mu.Unlock()

	if appConfig == nil {
		viper.AddConfigPath(path)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

		// Init config
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("[Config] - Failed to read config %v", err)
		}
		var cfg AppConfiguration

		if err := viper.Unmarshal(&cfg); err != nil {
			panic("unable to unmarshal config")
		}

		appConfig = &cfg
	}

	return appConfig
}
