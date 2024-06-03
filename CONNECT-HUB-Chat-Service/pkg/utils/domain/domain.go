package domain

import "time"

type Messages struct {
	FriendShipID string `gorm:"primaryKey;autoincrement;unique; type:integer"`
	Users        string
	Friend       string
	UpdateAt     time.Time
	Status       string `gorm:"default:pending"`
}
