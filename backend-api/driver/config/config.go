package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	confPathEnvName = "CONFIG_PATH"
)

type RawConfig struct {
	Config Config `yaml:"config"`
}

type Config struct {
	DBConfig       DBConfig       `yaml:"db"`
	SearchConfig   SearchConfig   `yaml:"search_engine"`
	CacheConfig    CacheConfig    `yaml:"cache"`
	EventBusConfig EventBusConfig `yaml:"event_bus"`
}

type dbDriver string

const (
	Pgx     dbDriver = "pgx"
	IndexES string   = "gallery"
)

type DBConfig struct {
	Driver  dbDriver `yaml:"driver"`
	ConnStr string   `yaml:"conn_str"`
}

type SearchConfig struct {
	Address string `yaml:"address"`
}

type CacheConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
}

type EventBusConfig struct {
	PublishAddress string `yaml:"publish_address"`
}

func Get() (cfg Config) {
	configPath := os.Getenv(confPathEnvName)
	log.Printf("getting config from %s\n", configPath)
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rawCfg := RawConfig{}
	err = yaml.NewDecoder(file).Decode(&rawCfg)
	if err != nil {
		panic(err)
	}
	return rawCfg.Config
}
