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

type ServerConfig struct {
	ServerPort string
}

type MysqlConfig struct {
	DSN string
}

type RedisConfig struct{}

func LoadEnv() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return &Config{}, nil
	}
	server_port := GetEnv("API_PORT")
	db_config := loadDBConfig() // 数据库配置
	config := &Config{
		ServerPort:  server_port,
		DSN:         db_config.buildDSN(),
		RedisConfig: RedisConfig{},
	}
	return config, nil
}

func GetEnv(env string) string {
	return os.Getenv(env)
}
