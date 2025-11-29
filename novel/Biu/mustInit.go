package Biu

import (
	"fmt"
	"log"
	"novel-launch/novel/middleware/db"
	models "novel-launch/novel/modells/db"
	"os"
)

func MustInit() error {
	err := initdb()
	if err != nil {
		return err
	}

	return nil
}

func initdb() error {
	user := getenv("DB_USER", "root")
	pass := getenv("DB_PASS", "")
	host := getenv("DB_HOST", "127.0.0.1")
	port := getenv("DB_PORT", "3306")
	name := getenv("DB_NAME", "novel_launch")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	log.Println("Dsn: ", dsn)
	if err := db.Init(dsn); err != nil {
		return err
	}
	if err := db.Get().AutoMigrate(&models.Books{}); err != nil {
		return err
	}
	log.Println("Database init Success")
	return nil
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
