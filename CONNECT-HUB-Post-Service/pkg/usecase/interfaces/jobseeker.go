package interfaces

import "github.com/ARunni/ConnetHub_post/pkg/utils/models"

type JobseekerPostUsecase interface {
	CreatePost(post models.CreatePostReq) (models.CreatePostRes, error)
	GetOnePost(postId int) (models.CreatePostRes, error)
	GetAllPost() (models.AllPostData, error)
	UpdatePost(post models.EditPostReq) (models.EditPostRes, error)
	DeletePost(postId, JobseekerId int) (bool, error)
	CreateCommentPost(postId, userId int, comment string) (bool, error)
	UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error)
	DeleteCommentPost(postId, userId, commentId int) (bool, error)
	AddLikePost(postId, userId int) (bool, error)
	RemoveLikePost(postId, userId int) (bool, error)
}
