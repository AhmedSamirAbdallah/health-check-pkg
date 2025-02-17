package healthcheck

import (
	"healthcheck/internal/redis"
	"time"
)

type RedisChecker struct {
	Addr     string
	Password string
	DB       int
	Key      string
	Val      string
}

func (r *RedisChecker) Name() string {
	return "redis"
}

func (r *RedisChecker) Check() map[string]interface{} {
	startTime := time.Now()

	redis.InitRedis(r.Addr, r.Password, r.DB)

	connection := redis.CheckRedisConnection()
	writeOK := redis.CheckWriteOnRedis(r.Key, r.Val)
	readOK := redis.CheckReadOnRedis(r.Key)

	return map[string]interface{}{
		"latency":    time.Since(startTime).String(),
		"connection": connection,
		"read":  readOK,
		"write": writeOK,
	}
}
