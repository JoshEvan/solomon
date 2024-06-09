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
	DBConfig DBConfig `yaml:"db"`
}

type dbDriver string

const (
	Pgx dbDriver = "pgx"
)

type DBConfig struct {
	Driver  dbDriver `yaml:"driver"`
	ConnStr string   `yaml:"conn_str"`
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
