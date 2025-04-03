package repository

import (
	"go-people-api/internal/config"
	"log"
	_ "os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Глобальная переменная для БД
var DB *sqlx.DB

// Инициализация БД
func Connect() {
	dsn := config.Get().DatabaseDSN // Получаем DSN из конфига

	var err error
	DB, err = sqlx.Connect("postgres", dsn) // Присваиваем в глобальную переменную

	if err != nil {
		log.Fatal("❌ Ошибка подключения к БД:", err)
	}

	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	log.Println("✅ Подключение к БД успешно")
}
