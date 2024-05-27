package models

import "time"

type CreatePostReq struct {
	JobseekerId int    `json:"jobseeker_id"`
	Title       string `gorm:"size:255;not null" json:"title"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Image       []byte `json:"image"`
}

type CreatePostRes struct {
	ID          int       `json:"id"`
	JobseekerId int       `json:"jobseeker_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type AllPost struct {
	Posts []CreatePostRes `json:"posts"`
}

type EditPostReq struct {
	JobseekerId int    `json:"jobseeker_id"`
	PostId      int    `json:"post_id"`
	Title       string `gorm:"size:255;not null" json:"title"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Image       []byte `json:"image"`
}

type EditPostRes struct {
	JobseekerId int       `json:"jobseeker_id"`
	PostId      int       `json:"post_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	ImageUrl    string    `json:"image_url"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
