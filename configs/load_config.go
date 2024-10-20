package configs

import (
	"encoding/json"
	"io"
	"os"
)

type DataBaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"dbname"`
	SSLMode  string `json:"sslMode"`
}

func LoadConfig(filepath string) (DataBaseConfig, error) {
	var config DataBaseConfig

	file, err := os.Open(filepath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
