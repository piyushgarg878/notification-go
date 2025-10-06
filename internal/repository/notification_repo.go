package repository

import (
	"context"
	"time"
	"notification-service-go/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository interface {
	Save(ctx context.Context, n *model.Notification) error
	GetAll(ctx context.Context) ([]model.Notification, error)
}

type MongoNotificationRepo struct {
	collection *mongo.Collection
}

func NewMongoNotificationRepo(db *mongo.Database) *MongoNotificationRepo {
	return &MongoNotificationRepo{
		collection: db.Collection("notifications"),
	}
}

func (r *MongoNotificationRepo) Save(ctx context.Context, n *model.Notification) error {
	_, err := r.collection.InsertOne(ctx, n)
	return err
}

func (r *MongoNotificationRepo) GetAll(ctx context.Context) ([]model.Notification, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var notifications []model.Notification
	if err := cursor.All(ctx, &notifications); err != nil {
		return nil, err
	}
	return notifications, nil
}