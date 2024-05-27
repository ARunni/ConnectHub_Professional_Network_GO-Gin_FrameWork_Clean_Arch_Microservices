package interfaces

import "ConnetHub_post/pkg/utils/models"

type JobseekerPostRepository interface {
	GetOnePost(postId int) (models.CreatePostRes, error)
	CreatePost(post models.CreatePostRes) (models.CreatePostRes, error)
	GetAllPost() (models.AllPost, error)
	UpdatePost(post models.EditPostRes) (models.EditPostRes, error)
	DeletePost(postId, JobseekerId int) (bool, error)
	IsPostExistByPostId(postId int) (bool, error)
	IsPostExistByUserId(userId int) (bool, error)
}
