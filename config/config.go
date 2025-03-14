package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config structure matching the YAML file
type Config struct {
	App struct {
		Name        string
		Environment string
		SecretKey   string
	}
	Server struct {
		Host string
		Port int
	}
	Database struct {
		Driver   string
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		SSLMode  string
	}
	Auth struct {
		JWTSecret   string
		TokenExpiry time.Duration
	}
	Logging struct {
		Level  string
		Format string
	}
	Cache struct {
		RedisHost     string
		RedisPort     int
		RedisPassword string
		TTL           int
	}
	ThirdParty struct {
		StripeAPIKey  string
		EmailProvider string
		SMTP          struct {
			Host     string
			Port     int
			User     string
			Password string
		}
	}
}

var AppConfig Config

// LoadConfig initializes Viper and loads the configuration file
func LoadConfig() {
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Look in the root directory

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal into AppConfig struct
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	// Convert token expiry duration from string to time.Duration
	tokenExpiryStr := viper.GetString("auth.token_expiry")
	AppConfig.Auth.TokenExpiry, _ = time.ParseDuration(tokenExpiryStr)
}
