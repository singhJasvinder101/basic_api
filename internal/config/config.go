package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {
	// means YAML me jo env: "production" hai, vo is Env field
	// me load hoga. when cfg.field_name = "" means seroalize nhi hua

	Env         string     `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string     `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HttpServer  HttpServer `yaml:"http_server"`
}

type HttpServer struct {
	Host string	
	Port int
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// check in flags of go run command
		flags := flag.String("config", "", "Path to config file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("CONFIG_PATH env variable or --config flag is required to load config")
		}
	}

	_, err := os.Stat(configPath); os.IsNotExist(err)
	if err != nil {
		log.Fatalf("Failed to stat config file %v", err)
	}

	var cfg Config

	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Failed to read config file %v", err)
	}

	return &cfg
}
