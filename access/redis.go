package access

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var (
	client *RedisClient
)

const (
	ip         = "192.168.110.201"
	port       = "6379"
	password   = ""
	userPrefix = "user"
)

func init() {
	client = new(RedisClient)
	client.initRedis(ip, port, password)
}

type RedisClient struct {
	handler *redis.Client
}

func (r *RedisClient) initRedis(ip, port, password string) {
	var err error
	r.handler, err = newRedisClient(ip, port, password)
	if err != nil {
		log.Fatalln("[CONNECT REDIS] ERROR:", err.Error())
	} else {
		log.Print("[CONNECT REDIS SUCCESS]")
	}
}

func GetRedisClient() *RedisClient {
	return client
}

func newRedisClient(ip, port, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     ip + ":" + port,
		Password: password,
		DB:       0, // use default DB
	})

	return client, client.Ping(context.Background()).Err()
}

func (r *RedisClient) IncrTimes(id string, times int64) (int64, error) {
	result, err := r.handler.IncrBy(context.Background(), userPrefix+id, times).Result()
	if err != nil {
		log.Fatalln("failed to IncrTimes. err:", err.Error())
		return 0, err
	}
	log.Print(result)
	return result, nil
}

func (r *RedisClient) DecrTimes(id string, times int64) (int64, error) {
	result, err := r.handler.DecrBy(context.Background(), userPrefix+id, times).Result()
	if err != nil {
		log.Fatalln("failed to DecrTimes. err:", err.Error())
		return 0, err
	}
	log.Print(result)
	return result, nil
}

func (r *RedisClient) ModifyTimes(id string, times int64) (int64, error) {
	err := r.handler.Set(context.Background(), userPrefix+id, times, 0).Err()
	if err != nil {
		log.Fatalln("failed to ModifyTimes. err:", err.Error())
		return 0, err
	}
	return 0, nil
}

func (r *RedisClient) GetTimes(id string) (int64, error) {
	result, err := r.handler.Get(context.Background(), userPrefix+id).Int64()
	if err != nil {
		log.Fatalln("failed to GetTimes. err:", err.Error())
		return 0, err
	}
	log.Print(result)
	return result, nil
}

func (r *RedisClient) DelTimes(id string) error {
	err := r.handler.Del(context.Background(), userPrefix+id).Err()
	if err != nil {
		log.Fatalln("failed to GetTimes. err:", err.Error())
		return err
	}
	return nil
}
