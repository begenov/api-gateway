package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Log                   LoggerConfig
	POSTGRES              PostgresConfig `required:"true"`
	Register_Service_Addr string         `required:"true"`
	Server                ServerConfig   `required:"true"`
}

type LoggerConfig struct {
	Level  string `json:"level" required:"true"`
	Format string `json:"format" required:"true"`
}

type PostgresConfig struct {
	Host     string `required:"true"`
	Port     int64  `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Database string `required:"true"`
}

type ServerConfig struct {
	HTTP_Port        string        `required:"true"`
	READ_TIMEOUT     time.Duration `required:"true"`
	WRITE_TIMEOUT    time.Duration `required:"true"`
	MAX_HEADER_BYTES int           `required:"true"`
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
