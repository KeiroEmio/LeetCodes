package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"os/signal"
	"time"
)
import "dataStructureDev/go-zero/RedisNX/LockerNX"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	go client1()
	go client2()
	<-ctx.Done()
}

func getLockcli() LockerNX.Locker {
	clientRedis := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
	})
	ttl := 5 * time.Second
	return LockerNX.NewRedisLocker(clientRedis, ttl)
}
func client1() {
	key := "my_key"
	clientRedis := getLockcli()
	clientRedis.Lock(key)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "client1加锁")
	defer clientRedis.UnLock(key)
	time.Sleep(20 * time.Second)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "client1解锁")
}

func client2() {
	key := "my_key"
	clientRedis := getLockcli()
	clientRedis.Lock(key)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "client2加锁")
	defer clientRedis.UnLock(key)
	time.Sleep(10 * time.Second)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "client2解锁")
}
