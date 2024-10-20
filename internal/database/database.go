package database

import (
	"context"
	"fmt"
	"log"

	"github.com/Flectere/system_of_crush/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

type Database struct {
	Pool *pgxpool.Pool
}

// Инициализация базы данных
func InitDB(config configs.DataBaseConfig) *Database {
	connString := fmt.Sprintf("user=%s dbname=%s host=%s port=%d sslmode=%s password=%s", config.Username, config.Database, config.Host, config.Port, config.SSLMode, config.Password)

	db, _ := pgxpool.New(context.Background(), connString)

	err := db.Ping(context.Background())
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных: ", err)
	}

	return &Database{Pool: db}
}

// Закрыть соединение с базой данных
func CloseDB() {
	db.Close()
}
