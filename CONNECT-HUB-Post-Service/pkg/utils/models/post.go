package models

import "time"

type Post struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	JobseekerId int       `json:"jobseeker_id"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Comment struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	PostID      uint      `gorm:"not null" json:"post_id"`
	Comment     string    `gorm:"type:text;not null" json:"comment"`
	JobseekerId uint      `gorm:"not null" json:"jobseeker_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Like struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	JobseekerId    uint      `gorm:"not null" json:"jobseeker_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

