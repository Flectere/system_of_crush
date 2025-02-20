package config

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config = &config{}

type DataBaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"dbname"`
	SSLMode  string `json:"sslMode"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	JWT  string `json:"jwt_key"`
	Salt string `json:"password_salt"`
}

type config struct {
	DataBaseConfig DataBaseConfig `json:"db_config"`
	ServerConfig   ServerConfig   `json:"server_config"`
}

func InitConfig(filepath string) error {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Файл .env не найден")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, Config)
	if err != nil {
		return err
	}

	Config.DataBaseConfig.Username = os.Getenv("POSTGRES_USER")
	Config.DataBaseConfig.Password = os.Getenv("POSTGRES_PASSWORD")
	Config.ServerConfig.Salt = os.Getenv("SALT")
	Config.ServerConfig.JWT = os.Getenv("JWT_KEY")

	return nil
}
