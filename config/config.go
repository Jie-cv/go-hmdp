package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig
	MysqlConfig
	RedisConfig
}

type ServerConfig struct{}

type MysqlConfig struct{}

type RedisConfig struct{}

func getEnv(env string) string {
	return os.Getenv(env)
}

func Load() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", nil
	}
	port := getEnv("API_PORT")
	return port, nil
}
