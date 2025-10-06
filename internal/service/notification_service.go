package service

import (
	"context"
	"fmt"
	"notification-service-go/internal/model"
	"notification-service-go/internal/repository"
	"notification-service-go/internal/utils"
	"time"
)

type NotificationService struct {
	repo repository.NotificationRepository
}

func NewNotificationService(repo repository.NotificationRepository) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) SendNotification(ctx context.Context, n *model.Notification) error {
	n.CreatedAt = time.Now()
	n.Status = "pending"

	var err error
	switch n.Type {
	case model.Email:
		err = utils.SendEmail(n.Recipient, n.Subject, n.Message)
	case model.SMS:
		err = utils.SendSMS(n.Recipient, n.Message)
	default:
		return fmt.Errorf("unsupported notification type: %s", n.Type)
	}

	if err != nil {
		n.Status = "failed"
	} else {
		n.Status = "sent"
	}

	return s.repo.Save(ctx, n)
}

func (s *NotificationService) GetNotifications(ctx context.Context) ([]model.Notification, error) {
	return s.repo.GetAll(ctx)
}