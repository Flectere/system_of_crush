package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

type DataBaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslMode"`
}

type Database struct {
	Pool *pgxpool.Pool
}

// Инициализация базы данных
func InitDB(config DataBaseConfig) *Database {
	//Создание конфига подключения к базе данных
	// cfg := DataBaseConfig{
	// 	Host:     config.Host,
	// 	Port:     config.Port,
	// 	Username: config.Username,
	// 	Password: config.Password,
	// 	Database: config.Database,
	// 	SSLMode:  config.SSLMode,
	// }

	connString := "user=postgres dbname=system_of_crush password= host=localhost port=5432 sslmode=disable"

	//Создание соедения с базой данных
	db, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	return &Database{Pool: db}

}

// Закрыть соединение с базой данных
func CloseDb() {
	db.Close()
}
