package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Mongo    MongoConfig    `mapstructure:"mongo"`
	Jwt      JwtConfig      `mapstructure:"jwt"`
}

type AppConfig struct {
	Port        string `mapstructure:"port"`
	Environment string `mapstructure:"environment"`
	LogLevel    string `mapstructure:"log_level"`
}

type DatabaseConfig struct {
	MasterDSN string   `mapstructure:"master_dsn"`
	Replicas  []string `mapstructure:"replicas"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type MongoConfig struct {
	URI      string `mapstructure:"uri"`
	Database string `mapstructure:"database"`
}

type JwtConfig struct {
	SecretKey string `mapstructure:"secret_key"`
}

var (
	config Config
	once   sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		env := viper.GetString("ENV")
		if env == "" {
			env = "dev"
		}

		viper.SetConfigName(env)
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config/env/")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("Unable to decode into struct, %v", err)
		}

		log.Printf("Loaded configuration for environment: %s", env)
	})

	return &config
}
