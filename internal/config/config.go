package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func LoadConf() {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("config path not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := godotenv.Read(configPath)
	if err != nil {
		log.Fatalf("can not read config file:%s", &err.Error())
	}

	return &cfg
}
