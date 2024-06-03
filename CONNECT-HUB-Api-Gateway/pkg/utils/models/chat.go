package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Messages struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SenderID       uint               `json:"sender_id" bson:"sender_id"`
	ChatID         primitive.ObjectID `json:"chat_id" bson:"chat_id"`
	Seen           bool               `json:"seen" bson:"seen"`
	Image          string             `json:"image" bson:"image"`
	MessageContent string             `json:"message_content" bson:"message_content"`
	Timestamp      time.Time          `json:"timestamp" bson:"timestamp"`
}

type TempMessage struct {
	SenderID    string
	RecipientID string `json:"RecipientID" validate:"required"`
	Content     string `json:"Content" validate:"required"`
	Timestamp   string `json:"TimeStamp" validate:"required"`
}

type ChatRequest struct {
	FriendID string `query:"FriendID" validate:"required"`
	Offset   string `query:"Offset" validate:"required"`
	Limit    string `query:"Limit" validate:"required"`
}

type Message struct {
	SenderID    string    `json:"SenderID" validate:"required"`
	RecipientID string    `json:"RecipientID" validate:"required"`
	Content     string    `json:"Content" validate:"required"`
	Timestamp   time.Time `json:"TimeStamp" validate:"required"`
}

type Pagination struct {
	Limit  string
	OffSet string
}
