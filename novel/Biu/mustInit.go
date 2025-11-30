package Biu

import (
	"log"
	"novel-launch/novel/middleware/db"
	redismw "novel-launch/novel/middleware/redis"
	"os"
)

func MustInit() error {
	err := initdb()
	if err != nil {
		return err
	}

	if err := initredis(); err != nil {
		return err
	}

	return nil
}

func initdb() error {
	if err := db.Init(); err != nil {
		return err
	}

	log.Println("Database init Success")
	return nil
}

func initredis() error {
	if err := redismw.Init(); err != nil {
		return nil
	}
	log.Println("Redis init Success")
	return nil
}

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
