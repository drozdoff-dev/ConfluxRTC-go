package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env string `yaml:"env" env-default:"local"`

	HTTPServer HTTPServerConfig `yaml:"http_server"`
}

type HTTPServerConfig struct {
	Address     string        `yaml:"address" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func New(log *slog.Logger) *Config {
	// 1. Загружаем .env (если есть)
	envPath := os.Getenv("ENV_PATH")
	if envPath != "" {
		if err := godotenv.Load(envPath); err != nil {
			log.Error("cannot load .env file", slog.String("path", envPath), slog.Any("error", err))
		} else {
			log.Info("env file loaded", slog.String("path", envPath))
		}
	}

	// 2. Создаем конфиг с дефолтными значениями
	var cfg Config

	// 3. Определяем путь к конфигу
	configDir := os.Getenv("CONFIG_DIR")
	if configDir == "" {
		configDir = "./configs" // дефолтный путь
	}
	configPath := fmt.Sprintf("%s/config.yaml", configDir)

	// 4. Загружаем YAML конфиг (если есть)
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Error("cannot read config", slog.String("path", configPath), slog.Any("error", err))
	} else {
		log.Info("config loaded", slog.String("path", configPath))
	}

	// 5. Применяем env-переменные поверх (если есть)
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Error("cannot read env vars", slog.Any("error", err))
	}

	return &cfg
}
