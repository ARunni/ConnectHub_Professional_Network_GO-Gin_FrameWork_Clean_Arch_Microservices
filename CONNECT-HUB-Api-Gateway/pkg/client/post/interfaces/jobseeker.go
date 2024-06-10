package interfaces

import "github.com/ARunni/connectHub_gateway/pkg/utils/models"

type JobseekerPostClient interface {
	CreatePost(post models.CreatePostReq) (models.CreatePostRes, error)
	GetOnePost(postId int) (models.CreatePostRes, error)
	GetAllPost() (models.AllPost, error)
	UpdatePost(post models.EditPostReq) (models.EditPostRes, error)
	DeletePost(postId, JobseekerId int) (bool, error)
	CreateCommentPost(postId, userId int, comment string) (bool, error)
	UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error)
	DeleteCommentPost(postId, userId, commentId int) (bool, error)
	AddLikePost(postId, userId int) (bool, error)
	RemoveLikePost(postId, userId int) (bool, error)
}
