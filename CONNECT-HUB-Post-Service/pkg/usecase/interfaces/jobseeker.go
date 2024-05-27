package interfaces

import "ConnetHub_post/pkg/utils/models"

type JobseekerPostUsecase interface {
	CreatePost(post models.CreatePostReq) (models.CreatePostRes, error)
	GetOnePost(postId int) (models.CreatePostRes, error)
	GetAllPost() (models.AllPost, error)
	UpdatePost(post models.EditPostReq) (models.EditPostRes, error)
	DeletePost(postId, JobseekerId int) (bool, error)
}
