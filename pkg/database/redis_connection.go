package database

import (
	"context"
	"experiment/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisConnection struct {
	Client *redis.Client
	Prefix string
}

func NewRedisConnection(cfg config.RedisConfig) (*RedisConnection, error) {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: "",
		DB:       cfg.DB,
	}

	client := redis.NewClient(options)

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisConnection{Client: client, Prefix: cfg.Prefix}, nil
}

func (conn *RedisConnection) SetWithPrefix(key string, value interface{}) error {
	fullKey := fmt.Sprintf("%s:%s", conn.Prefix, key)
	return conn.Client.Set(conn.Client.Context(), fullKey, value, 0).Err()
}

// Put сохраняет значение в Redis с указанным ключом и TTL (временем жизни) с учетом префикса приложения
func (conn *RedisConnection) Put(key string, value interface{}, ttl int64) error {
	fullKey := fmt.Sprintf("%s:%s", conn.Prefix, key)
	return conn.Client.Set(context.Background(), fullKey, value, time.Duration(ttl)*time.Second).Err()
}

// Get возвращает значение из Redis по указанному ключу с учетом префикса приложения
func (conn *RedisConnection) Get(key string) (string, error) {
	fullKey := fmt.Sprintf("%s:%s", conn.Prefix, key)
	val, err := conn.Client.Get(context.Background(), fullKey).Result()

	// Проверка наличия ключа в Redis
	if err != nil {
		if err == redis.Nil {
			// Ключа нет в Redis
			return "", fmt.Errorf("key not found in Redis: %w", err)
		}
		// Произошла другая ошибка при получении значения
		return "", fmt.Errorf("error getting value from Redis: %w", err)
	}

	// Значение успешно получено из Redis
	return val, nil
}
