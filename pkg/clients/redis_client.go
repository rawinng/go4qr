package clients

import (
	"time"

	"github.com/go-redis/redis"
)

// create struct for Redis Host and Port
type RedisClient struct {
	Host     string
	Port     string
	instance *redis.Client
}

var singletonRedisClient *RedisClient

func GetRedisInstance() *RedisClient {
	if singletonRedisClient == nil {
		singletonRedisClient = &RedisClient{
			Host: "localhost",
			Port: "6379",
		}
		singletonRedisClient.createRedisClient()
	}
	return singletonRedisClient
}

// Create redis client from RedisConfig Host and Port
func (r *RedisClient) createRedisClient() {
	r.instance = redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		Password: "",
		DB:       0,
	})
}

// generate functions for Set, Get, SetWithExpire, Delete function for RedisClient
func (r *RedisClient) Set(key string, value interface{}) error {
	err := r.instance.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.instance.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisClient) SetWithExpire(key string, value interface{}, expirationSecond int) error {
	err := r.instance.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	err = r.instance.Expire(key, time.Duration(expirationSecond*int(time.Second))).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Delete(key string) error {
	err := r.instance.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
