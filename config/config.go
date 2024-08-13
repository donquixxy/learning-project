package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type JWTConfig struct {
	Issuer        string `mapstructure:"ISSUER"`
	Secret        string `mapstructure:"SECRET"`
	SecretRefresh string `mapstructure:"SECRET_REFRESH"`
}

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

type RabbitConfig struct {
	RabbitHost string `mapstructure:"RABBIT_HOST"`
	RabbitPort int    `mapstructure:"RABBIT_PORT"`
	Username   string `mapstructure:"RABBIT_USER"`
	Password   string `mapstructure:"RABBIT_PASS"`
}

var appConfig *AppConfiguration
var dbConfig *DatabaseConfig
var jwtConfig *JWTConfig
var rabbitConfig *RabbitConfig
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

func GetJwtConfig() *JWTConfig {
	mu.Lock()

	defer mu.Unlock()

	if jwtConfig == nil {
		viper.AddConfigPath(path)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

		// Init config
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("[Config] - Failed to read config %v", err)
		}
		var cfg JWTConfig

		if err := viper.Unmarshal(&cfg); err != nil {
			panic("unable to unmarshal config")
		}

		jwtConfig = &cfg
	}

	return jwtConfig
}

func GetRabbitConfig() *RabbitConfig {
	mu.Lock()
	defer mu.Unlock()

	if rabbitConfig == nil {
		viper.AddConfigPath(path)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("[Config] - Failed to read config %v", err)
		}

		var cfg RabbitConfig
		if err := viper.Unmarshal(&cfg); err != nil {
			panic("unable to unmarshal config")
		}

		rabbitConfig = &cfg
		return rabbitConfig
	}

	return rabbitConfig
}
