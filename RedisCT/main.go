package main

import (
	"context"
	"fmt"
	"time"
)

var ctx = context.Background()

func main() {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 地址
		Password: "",               // Redis 密码
		DB:       0,                // Redis DB
	})

	key := "user:id:99999" // 一个不存在的 key

	// 模拟缓存穿透
	for i := 0; i < 100; i++ {
		value, err := rdb.Get(ctx, key).Result()
		if err == redis.Nil {
			fmt.Println("Cache miss, accessing database for key:", key)
			// 模拟数据库访问
			time.Sleep(10 * time.Millisecond) // 模拟查询数据库的延迟

			// 假设数据库返回了空值，我们不缓存它
			// 继续下一次查询
		} else if err != nil {
			fmt.Println("Error accessing Redis:", err)
			return
		} else {
			fmt.Println("Cache hit, value:", value)
		}
	}
}
