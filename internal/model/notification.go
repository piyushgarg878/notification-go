package model

import "time"

type NotificationType string

const (
	Email NotificationType = "email"
	SMS   NotificationType = "sms"
)

type Notification struct {
	ID        string           `json:"id" bson:"_id,omitempty"`
	UserID    string           `json:"user_id" bson:"user_id"`
	Type      NotificationType `json:"type" bson:"type"`
	Recipient string           `json:"recipient" bson:"recipient"`
	Subject   string           `json:"subject,omitempty" bson:"subject,omitempty"`
	Message   string           `json:"message" bson:"message"`
	Status    string           `json:"status" bson:"status"`
	CreatedAt time.Time        `json:"created_at" bson:"created_at"`
}