package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	App         AppConfig
	DB          DBConfig
	Redis       RedisConfig
	RedisCache  RedisConfig
	FileStorage FileStorageConfig
}

type AppConfig struct {
	Name  string
	Env   string
	Debug bool
	Host  string
	Port  int
}

type DBConfig struct {
	Connection string //mysql, postgresql, sqlite ...
	Host       string
	Port       int
	Database   string
	Username   string
	Password   string
}

type RedisConfig struct {
	Host   string
	Port   int
	DB     int
	Prefix string
}

type RedisCacheConfig struct {
	Host    string
	Port    int
	CacheDB int
	Prefix  string
}

type FileStorageConfig struct {
	FileSystemDisk   string
	AwsEndpoint      string
	AwsDefaultRegion string
	AwsBucket        string
	AwsKey           string
	AwsSecret        string
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func Init() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %w", err)
	}

	app := AppConfig{
		Name:  getEnv("APP_NAME", "Golang"),
		Env:   getEnv("APP_ENV", "Local"),
		Debug: getEnvAsBool("APP_DEBUG", true),
		Host:  getEnv("APP_HOST", "localhost"),
		Port:  getEnvAsInt("APP_PORT", 8080),
	}

	db := DBConfig{
		Connection: getEnv("DB_CONNECTION", "mysql"),
		Host:       getEnv("DB_HOST", "localhost"),
		Port:       getEnvAsInt("DB_PORT", 3306),
		Database:   getEnv("DB_DATABASE", ""),
		Username:   getEnv("DB_USERNAME", ""),
		Password:   getEnv("DB_PASSWORD", ""),
	}

	redis := RedisConfig{
		Host:   getEnv("REDIS_HOST", ""),
		Port:   getEnvAsInt("REDIS_PORT", 6379),
		DB:     getEnvAsInt("REDIS_DB", 0),
		Prefix: getEnv("REDIS_PREFIX", ""),
	}

	redisCache := RedisConfig{
		Host:   getEnv("REDIS_HOST", ""),
		Port:   getEnvAsInt("REDIS_PORT", 6379),
		DB:     getEnvAsInt("REDIS_CACHE_DB", 0),
		Prefix: getEnv("REDIS_PREFIX", ""),
	}

	fileStorage := FileStorageConfig{
		FileSystemDisk:   getEnv("FILESYSTEM_DISK", "storage"),
		AwsEndpoint:      getEnv("AWS_ENDPOINT", ""),
		AwsBucket:        getEnv("AWS_BUCKET", ""),
		AwsDefaultRegion: getEnv("AWS_DEFAULT_REGION", ""),
		AwsKey:           getEnv("AWS_KEY", ""),
		AwsSecret:        getEnv("AWS_SECRET", ""),
	}

	config := &Config{
		App:         app,
		DB:          db,
		Redis:       redis,
		RedisCache:  redisCache,
		FileStorage: fileStorage,
	}

	return config, nil
}
