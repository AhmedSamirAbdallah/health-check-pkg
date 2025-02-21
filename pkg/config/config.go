package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceName    string
	DBURI          string
	DatabaseName   string
	CollectionName string
	KafkaBroker    string
	KafkaTopic     string
	KafkaGroupID   string
	RedisURL       string
	RedisPassword  string
	RedisDB        int
	RedisKey       string
	RedisVal       string
	TemporalUrl    string
	WithTLS        bool
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Printf("Error loading environment file: %v\n", err)
	}
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Printf("Invalid value for REDIS_DB: %v. Using default value %d.\n", err, redisDB)
		redisDB = 0
	}

	return &Config{
		ServiceName:    os.Getenv("SERVICE_NAME"),
		DBURI:          os.Getenv("DB_URI"),
		DatabaseName:   os.Getenv("DATABASE_NAME"),
		CollectionName: os.Getenv("COLLECTION_NAME"),
		KafkaBroker:    os.Getenv("KAFKA_BROKER"),
		KafkaTopic:     os.Getenv("KAFKA_TOPIC"),
		KafkaGroupID:   os.Getenv("KAFKA_GROUP_ID"),
		RedisURL:       os.Getenv("REDIS_URL"),
		RedisPassword:  os.Getenv("REDIS_PASSWORD"),
		RedisDB:        redisDB,
		RedisKey:       os.Getenv("REDIS_KEY"),
		RedisVal:       os.Getenv("REDIS_VALUE"),
		TemporalUrl:    os.Getenv("TEMPORAL_URL"),
		WithTLS:        os.Getenv("TEMPORAL_WITH_TLS") == "true",
	}, nil
}
