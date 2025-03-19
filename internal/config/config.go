package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Host string `yaml:"host" env:"HOST" env-default:"localhost" validate:"required"`
	Port string `yaml:"port" env:"PORT" env-default:"8080" validate:"required,numeric"`
}

type DBConfig struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost" validate:"required"`
	Port     int    `yaml:"port" env:"DB_PORT" env-default:"5432" validate:"required,numeric"`
	User     string `yaml:"user" env:"DB_USER" env-default:"user" validate:"required"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-default:"password" validate:"required"`
	Name     string `yaml:"name" env:"DB_NAME" env-default:"marketplace" validate:"required"`
}

type Config struct {
	Env        string     `yaml:"env" env:"ENV" env-default:"prod" validate:"oneof=dev prod test"`
	HTTPServer HTTPServer `yaml:"http_server" validate:"required"`
	DB         DBConfig   `yaml:"db" validate:"required"`
}

func New() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flag.StringVar(&configPath, "config", "./config/local.yaml", "path to config file")
		flag.Parse()
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &cfg, nil
}
