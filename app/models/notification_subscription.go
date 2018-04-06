package models

import (
	"time"
)

type NotificationSubscription struct {
	Id           uint      `gorm:"primary_key" json:"id"`
	SubscriberId int       `json:"subscriber_id"`
	PublisherId  int       `json:"publisher_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	// Plus add an index on subscriber and publisher.
}

type NotificationSubscriptionJSON struct {
	Type string `json:"type"`
	NotificationSubscription
}

func NewNotificationSubscriptionJSON(ns NotificationSubscription) NotificationSubscriptionJSON {
	nsJSON := NotificationSubscriptionJSON{"notification_subsription", ns}
	return nsJSON
}
