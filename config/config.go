package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug bool         `yaml:"is_debug"`
	Listen  ListenConfig `yaml:"listen"`
	Logger  LoggerConfig `yaml:"logging"`
}

type ListenConfig struct {
	Type string `yaml:"type"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type LoggerConfig struct {
	Writer string `yaml:"writer"`
	Level  string `yaml:"level"`
}

func New() (Config, error) {
	instance := Config{}
	log.Println("Reading application's configuration")

	useEnvConfig := os.Getenv("USE_ENV_CONFIG") == "true"

	if !useEnvConfig {
		if err := cleanenv.ReadConfig("config.yml", &instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println("Configuration file error:")
			log.Println(help)
			return Config{}, err
		}
	}

	if err := cleanenv.ReadEnv(&instance); err != nil {
		return Config{}, err
	}

	return instance, nil
}
