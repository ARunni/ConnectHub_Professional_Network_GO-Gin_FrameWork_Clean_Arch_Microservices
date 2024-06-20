package models

import "time"

type NotificationReq struct {
	UserID     int       `json:"user_id"`
	SenderID   int       `json:"sender_id"`
	SenderName string    `json:"sender_name"`
	PostID     int       `json:"post_id"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type NotificationResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id" gorm:"column:sender_id"`
	Username  string `json:"username"`
	PostID    int    `json:"post_id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type AllNotificationResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id" gorm:"column:sender_id"`
	Username  string `json:"username"`
	PostID    int    `json:"post_id"`
	Message   string `json:"message"`
	Read      bool   `json:"read"`
	CreatedAt string `json:"created_at"`
}
type Notification struct {
	ID         int       `json:"id" gorm:"primaryKey"`
	UserID     int       `json:"user_id"`
	SenderID   int       `json:"sender_id"`
	SenderName string    `json:"sender_name"`
	PostID     int       `json:"post_id"`
	Message    string    `json:"message"`
	Read       bool      `json:"read" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserData struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}
