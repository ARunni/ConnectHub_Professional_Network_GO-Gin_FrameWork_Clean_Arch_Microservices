package models

import (
	"time"
)

type Message struct {
	ID          string    `bson:"_id"`
	SenderID    string    `bson:"senderid"`
	RecipientID string    `bson:"recipientid"`
	Content     string    `bson:"content"`
	Timestamp   time.Time `bson:"timestamp"`
}

type MessageReq struct {
	SenderID    string    `bson:"senderid"`
	RecipientID string    `bson:"recipientid"`
	Content     string    `bson:"content"`
	Timestamp   time.Time `bson:"timestamp"`
}

type Pagination struct {
	Limit  string
	OffSet string
}
