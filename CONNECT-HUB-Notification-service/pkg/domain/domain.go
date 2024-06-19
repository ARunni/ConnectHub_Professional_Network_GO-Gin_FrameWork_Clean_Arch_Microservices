package domain

import "time"

type Notification struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	SenderID  int       `json:"sender_id"`
	PostID    int       `json:"post_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
