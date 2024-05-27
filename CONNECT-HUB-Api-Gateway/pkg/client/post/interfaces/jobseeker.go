package interfaces

import "connectHub_gateway/pkg/utils/models"

type JobseekerPostClient interface {
	CreatePost(post models.CreatePostReq) (models.CreatePostRes, error)
	GetOnePost(postId int) (models.CreatePostRes, error)
	GetAllPost() (models.AllPost, error)
	UpdatePost(post models.EditPostReq) (models.EditPostRes, error)
	DeletePost(postId, JobseekerId int) (bool, error)
}

