package models

import "time"

type CreatePostReq struct {
	JobseekerId int    `json:"jobseeker_id"`
	Title       string `gorm:"size:255;not null" json:"title"`
	Content     string `gorm:"type:text;not null" json:"content"`
	Image       []byte `json:"image"`
}

type CreatePostRes struct {
	ID          int         `json:"id"`
	JobseekerId int         `json:"jobseeker_id"`
	Title       string      `gorm:"size:255;not null" json:"title"`
	Content     string      `gorm:"type:text;not null" json:"content"`
	ImageUrl    string      `json:"image_url"`
	Comments    []CommentData `gorm:"default:null" json:"comments"`
	Likes       int         `gorm:"default:0" json:"likes"`
	CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"`
}

type CreatePostResp struct {
	ID          int         `json:"id"`
	JobseekerId int         `json:"jobseeker_id"`
	Title       string      `gorm:"size:255;not null" json:"title"`
	Content     string      `gorm:"type:text;not null" json:"content"`
	ImageUrl    string      `json:"image_url"`
	CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"`
}

type CommentData struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Comment     string    `gorm:"type:text;not null" json:"comment"`
	JobseekerId uint      `gorm:"not null" json:"jobseeker_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type AllPost struct {
	Posts []CreatePostResp `json:"posts"`
}
type AllPostData struct {
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
	PostId      uint      `json:"post_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	ImageUrl    string    `json:"image_url"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserData struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}