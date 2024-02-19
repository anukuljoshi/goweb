package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port int    `env:"PORT" envDefault:"8000"`
	Env  string `env:"ENV,required"`
	DSN  string `env:"DSN,required"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
