package redis

import (
	"context"
	"novel-launch/novel/resource"
	"os"
	"strconv"

	goredis "github.com/redis/go-redis/v9"
)

func Init() error {
	addr := getenv("REDIS_ADDR", "127.0.0.1:6379")
	pass := getenv("REDIS_PASS", "")
	dbstr := getenv("REDIS_DB", "11")
	n, _ := strconv.Atoi(dbstr)
	c := goredis.NewClient(&goredis.Options{Addr: addr, Password: pass, DB: n})
	if err := c.Ping(context.Background()).Err(); err != nil {
		return err
	}
	resource.Rdscli = c
	return nil
}

func Get() *goredis.Client { return resource.Rdscli }

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
