package models

import "time"

type JobSeekerLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=6,max=20"`
}

type JobSeekerSignUp struct {
	Email           string `json:"email" binding:"required" validate:"required,email"`
	Password        string `json:"password" binding:"required" validate:"min=6,max=20"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"min=6,max=20"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	DateOfBirth     string `json:"date_of_birth"`
	Gender          string `json:"gender"`
}

type JobSeekerDetailsResponse struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

type TokenJobSeeker struct {
	JobSeeker JobSeekerDetailsResponse
	Token     string
}

type JobSeekerProfile struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

type JobSeekerProfileReq struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
}

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
type CommentData struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Comment     string    `gorm:"type:text;not null" json:"comment"`
	JobseekerId uint      `gorm:"not null" json:"jobseeker_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
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

type CreateCommentPost struct {
	PostId  int    `json:"post_id"`
	UserId  int    `json:"user_id"`
	Comment string `json:"comment"`
}

type CreateCommentPostReq struct {
	PostId  int    `json:"post_id"`
	UserId  int    `json:"user_id"`
	Comment string `json:"comment"`
}
type UpdateCommentPost struct {
	CommentId int    `json:"comment_id"`
	Comment  string `json:"comment"`
	PostId  int    `json:"post_id"`
	UserId  int    `json:"user_id"`
}

type UpdateCommentPostReq struct {
	CommentId int    `json:"comment_id"`
	Comment  string `json:"comment"`
	PostId  int    `json:"post_id"`
	UserId  int    `json:"user_id"`
}

type DeleteCommentPost struct {
	CommentId int `json:"comment_id"`
	PostId  int `json:"post_id"`
	UserId  int `json:"user_id"`

}

type DeleteCommentPostReq struct {
	CommentId int `json:"comment_id"`
	PostId  int `json:"post_id"`
	UserId  int `json:"user_id"`

}
