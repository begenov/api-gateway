package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Log                 LoggerConfig
	POSTGRES            PostgresConfig `required:"true"`
	RegisterServiceAddr string         `required:"true"`
}

type LoggerConfig struct {
	Level  string `json:"level" required:"true"`
	Format string `json:"format" required:"true"`
	HTTP   HTTPServerConfig
}

type PostgresConfig struct {
	Host     string `required:"true"`
	Port     int64  `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Database string `required:"true"`
}

type HTTPServerConfig struct {
	Port           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func NewConfig() *Config {
	envFile := "./.env"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}
	var config Config
	err = envconfig.Process("", &config)
	if err != nil {
		log.Fatalln(err)
	}
	return &config
}
