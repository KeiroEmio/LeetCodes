package LockerNX

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Locker interface {
	Lock(key string)
	UnLock(key string)
}

type redisLocker struct {
	client redis.UniversalClient
	llt    time.Duration
	done   chan struct{}
}

func NewRedisLocker(client redis.UniversalClient, llt time.Duration) Locker {
	return &redisLocker{
		client: client,
		llt:    llt,
	}
}

func (r *redisLocker) Lock(key string) {
	for {
		ok, err := r.client.SetNX(context.Background(), key, "", r.llt).Result()
		if err != nil || !ok {
			time.Sleep(100 * time.Second)
			continue
		}
		break
	}
	r.done = make(chan struct{})
	go func() {
		for {
			select {
			case <-r.done:
				fmt.Println("程序退出，释放锁")
				return
			default:
				_, err := r.client.Expire(context.Background(), key, r.llt).Result()
				if err != nil {
					log.Print("err:", err)
				}
				<-time.After(r.llt / 2)
			}
		}
	}()
}

func (r *redisLocker) UnLock(key string) {
	defer close(r.done)
	err := r.client.Del(context.Background(), key).Err()
	if err != nil {
		log.Println("unlock err:", err)
	}
	log.Println("unlock 解锁成功")
}
