package models

import "time"

type Notification struct {
	UserID    int       `json:"user_id"`
	SenderID  int       `json:"sender_id"`
	JobID    int       `json:"job_id"`
	Message   string    `json:"Message"`
	CreatedAt time.Time `json:"created_at"`
}