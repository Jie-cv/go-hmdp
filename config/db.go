package config

import "fmt"

type DBConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	DBName   string
	Charset  string
}

func loadDBConfig() *DBConfig {
	return &DBConfig{
		UserName: GetEnv("MYSQL_USER"),
		Password: GetEnv("MYSQL_PASSWORD"),
		Host:     GetEnv("MYSQL_HOST"),
		Port:     GetEnv("MYSQL_PORT"),
		DBName:   GetEnv("MYSQL_DB_NAME"),
		Charset:  GetEnv("MYSQL_CHARSET"),
	}
}

func (d *DBConfig) buildDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		d.UserName, d.Password, d.Host, d.Port, d.DBName, d.Charset)
}
