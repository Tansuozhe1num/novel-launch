package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init opens a GORM MySQL connection. If dsn is empty, it builds one from env.
func Init(dsn string) error {
	if dsn == "" {
		user := getenv("DB_USER", "root")
		pass := getenv("DB_PASS", "")
		host := getenv("DB_HOST", "127.0.0.1")
		port := getenv("DB_PORT", "3306")
		name := getenv("DB_NAME", "novel_launch")
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	}
	g, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Warn)})
	if err != nil {
		return err
	}
	DB = g
	return nil
}

func Get() *gorm.DB { return DB }

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
