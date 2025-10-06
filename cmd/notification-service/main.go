package main

import (
	"context"
	"log"
	"notification-service-go/internal/config"
	"notification-service-go/internal/handler"
	"notification-service-go/internal/repository"
	"notification-service-go/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.LoadConfig()

	// MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal("❌ Failed to create MongoDB client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}

	db := client.Database(cfg.Database)
	repo := repository.NewMongoNotificationRepo(db)
	svc := service.NewNotificationService(repo)
	h := handler.NewNotificationHandler(svc)

	r := gin.Default()

	r.POST("/notify", h.SendNotification)
	r.GET("/notifications", h.GetAllNotifications)

	log.Println("🚀 Notification Service running on",cfg.Port)
	if err := r.Run(cfg.Port); err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
