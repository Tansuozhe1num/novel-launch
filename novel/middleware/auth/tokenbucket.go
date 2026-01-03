package auth

import (
	"context"
	"fmt"
	redis2 "github.com/redis/go-redis/v9"
	constance "novel-launch/novel/Biu/const"
	"novel-launch/novel/resource"
	"time"
)

var tokenBucketLuaScript = `
-- KEYS[1]  限流 key
-- ARGV[1]  当前时间戳（毫秒）
-- ARGV[2]  令牌生成速率（token / 秒）
-- ARGV[3]  桶容量
-- ARGV[4]  本次请求消耗的令牌数

local key = KEYS[1]
local now = tonumber(ARGV[1])
local rate = tonumber(ARGV[2])
local capacity = tonumber(ARGV[3])
local cost = tonumber(ARGV[4])

local res = redis.call('HMGET', key, 'tokens', 'timestamp')
local tokens = res[1]
local last_time = res[2]

if tokens == false or last_time == false then
    tokens = capacity
    last_time = now
else
    tokens = tonumber(tokens)
    last_time = tonumber(last_time)
end

-- 毫秒 → 秒
local delta = math.max(0, now - last_time) / 1000
local new_tokens = math.min(capacity, tokens + delta * rate)

if new_tokens < cost then
    redis.call('HMSET', key,
        'tokens', new_tokens,
        'timestamp', now
    )
    return 0
else
    redis.call('HMSET', key,
        'tokens', new_tokens - cost,
        'timestamp', now
    )
    return 1
end
`

// AllowRequest 令牌桶限流
func AllowRequest(ctx context.Context, key string, rate float64, capacity int) (bool, error) {
	if key == "" {
		return false, nil
	}
	if rate < 0 {
		fmt.Println("warning: token rate < 0")
		rate = constance.TokenRate
	}
	if capacity < 0 || capacity > 1000000 {
		fmt.Println("warning: token capacity must be reset")
		capacity = constance.TokenCapacity
	}

	tokenBucketLua := redis2.NewScript(tokenBucketLuaScript)
	res, err := tokenBucketLua.Run(
		ctx,
		resource.Rdscli,
		[]string{key},
		time.Now().UnixMilli(),
		rate, // token / 秒
		capacity,
		1,
	).Int64()

	if err != nil {
		return false, err
	}
	return res == 1, nil

}
