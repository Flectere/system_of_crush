package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Flectere/system_of_crush/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

type Database struct {
	Pool *pgxpool.Pool
}

// Инициализация базы данных
func InitDB(config config.DataBaseConfig) *Database {
	connString := fmt.Sprintf("user=%s dbname=%s host=%s port=%d sslmode=%s password=%s", config.Username, config.Database, config.Host, config.Port, config.SSLMode, config.Password)

	db, _ := pgxpool.New(context.Background(), connString)

	var err error
	for i := 1; i <= 10; i++ {
		log.Printf("Попытка подключения к базе данных: %v\n", i)

		err = db.Ping(context.Background())
		if err == nil {
			log.Println("Подключение к базе данных успешно!")
			break
		}
		time.Sleep(time.Second * 10)
	}

	if err != nil {
		log.Fatal("Неудалось подключиться к базе данных: ", err)
	}

	return &Database{Pool: db}
}

// Закрыть соединение с базой данных
func CloseDB() {
	db.Close()
}
