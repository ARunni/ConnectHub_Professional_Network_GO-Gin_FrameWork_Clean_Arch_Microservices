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
	CreateCommentPost(postId, userId int, comment string) (bool, error)
	IsCommentIdExist(commentId int) (bool, error)
	IsCommentIdBelongsUserId(commentId, userId int) (bool, error)
	UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error)
	DeleteCommentPost(postId, userId, commentId int) (bool, error)
	IsLikeExist(postId, userId int) (bool, error)
	AddLikePost(postId, userId int) (bool, error)
	RemoveLikePost(postId, userId int) (bool, error)
}
