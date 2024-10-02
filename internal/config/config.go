package config

import "github.com/caarlos0/env/v11"

type Config struct {
	MysqlURL   string `env:"MYSQL_URL" envDefault:"user:password@tcp(localhost:3306)/db"`
	ServerPort string `env:"SERVER_PORT" envDefault:":8081"`
	SourceURL  string `env:"SOURCE_URL" envDefault:"https://api.nbrb.by/exrates/rates?periodicity=0"`
}

func New() (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
