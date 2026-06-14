package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `toml:"env" env-default:"prod"`
	StoragePath string        `toml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `toml:"token_ttl"  env-required:"true"`
	GRPC        GRPCConfig    `toml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `toml:"port"`
	Timeout time.Duration `toml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("Путь к конфигу не может быть пустым")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Конфиг-файла не существует: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Ошибка чтения конфига: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	// --config="path/to/config.toml"
	flag.StringVar(&res, "config", "", "Путь до конфиг-файла")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
