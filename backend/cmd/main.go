package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/handler"
	"github.com/klausfun/Augventure/pkg/infrastructure"
	"github.com/klausfun/Augventure/pkg/repository"
	"github.com/klausfun/Augventure/pkg/service"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.hostAndPort"),
		Password: "", // os.Getenv("REDIS_PASSWORD")
		DB:       viper.GetInt("redis.dbname"),
	})

	ping, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		logrus.Fatalf("failed to initialize redis: %s", err.Error())
	}
	fmt.Println(ping)

	// Инициализация Yandex Object Storage
	cloudStorage, err := infrastructure.NewS3Storage(
		viper.GetString("s3.bucketName"),
		viper.GetString("s3.region"),
		viper.GetString("s3.endpoint"),
		os.Getenv("S3_ACCESS_KEY"),
		os.Getenv("S3_SECRET_KEY"),
	)
	if err != nil {
		logrus.Fatalf("failed to initialize cloud storage: %s", err.Error())
	}

	repos := repository.NewRepository(db, redisClient)
	services := service.NewService(repos, cloudStorage)
	handlers := handler.NewHandler(services)

	srv := new(augventure.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
