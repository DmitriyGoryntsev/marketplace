package config

import (
	"flag"
	"os"

	"github.com/DmitriyGoryntsev/marketplace/pkg/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

type HTTPServer struct {
	Host string `yaml:"host" env:"HOST" env-required:"true" env-default:"localhost"`
	Port string `yaml:"port" env:"PORT" env-required:"true" env-default:"8080"`
}
type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true" env-default:"prod"`

	HTTPServer HTTPServer `yaml:"http_server"`

	DBConfig postgres.DBConfig `yaml:"db"`
}

func New() (*Config, error) {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			zap.L().Fatal("config path is empty")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		zap.L().Fatal("config file does not exist")
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		zap.L().Fatal("failed to read config", zap.Error(err))
	}

	return &cfg, nil
}
