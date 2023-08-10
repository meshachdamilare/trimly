package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {

	// Database Configuration
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	// Server Port
	ServerPort string `mapstructure:"SERVER_PORT"`

	// JWT Configuration
	Secret_Key string `mapstructure:"SECRET_KEY"`

	//Email Configuration
	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`

	Origin string `mapstructure:"CLIENT_ORIGIN"`
}

var (
	Config *Configuration
)

func Setup() {
	Config = &Configuration{}

	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable reading settings file, %s", err)
	}

	err := viper.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}

func GetConfig() *Configuration {
	return Config
}
