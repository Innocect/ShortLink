package main

import (
	"testing"

	"bou.ke/monkey"
	"github.com/ashu/redisDao"
	"github.com/go-redis/redis"
)

func TestMain_shouldReturnPanic_WhenRedisFailsToInit(t *testing.T) {
	setUpPatchForRedisMock()
	defer monkey.UnpatchAll()

	main()
}

func setUpPatchForRedisMock() {
	monkey.Patch(redisDao.RedisConnection, func() *redis.Client {
		return nil
	})
}
