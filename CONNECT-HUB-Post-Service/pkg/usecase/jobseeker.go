package usecase

import (
	"ConnetHub_post/pkg/config"
	"ConnetHub_post/pkg/helper"
	repo "ConnetHub_post/pkg/repository/interfaces"
	"ConnetHub_post/pkg/usecase/interfaces"
	"ConnetHub_post/pkg/utils/models"
	"errors"
)

type jobseekerJobUseCase struct {
	postRepository repo.JobseekerPostRepository
}

func NewJobseekerpostUseCase(repo repo.JobseekerPostRepository) interfaces.JobseekerPostUsecase {
	return &jobseekerJobUseCase{
		postRepository: repo,
	}
}

func (ju *jobseekerJobUseCase) CreatePost(post models.CreatePostReq) (models.CreatePostRes, error) {

	var postReq models.CreatePostRes

	cfg, err := config.LoadConfig()
	if err != nil {
		return models.CreatePostRes{}, err
	}

	h := helper.NewHelper(cfg)
	url, err := h.AddImageToAwsS3(post.Image)
	if err != nil {
		return models.CreatePostRes{}, err
	}

	postReq.Content = post.Content
	postReq.Title = post.Title
	postReq.JobseekerId = post.JobseekerId
	postReq.ImageUrl = url

	postData, err := ju.postRepository.CreatePost(postReq)

	if err != nil {
		return models.CreatePostRes{}, err
	}
	return postData, nil

}

func (ju *jobseekerJobUseCase) GetOnePost(postId int) (models.CreatePostRes, error) {

	ok, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		return models.CreatePostRes{}, err
	}

	if !ok {
		return models.CreatePostRes{}, errors.New("post does not exist")
	}
	postData, err := ju.postRepository.GetOnePost(postId)
	if err != nil {
		return models.CreatePostRes{}, err
	}
	return postData, nil
}

func (ju *jobseekerJobUseCase) GetAllPost() (models.AllPost, error) {
	postData, err := ju.postRepository.GetAllPost()

	if err != nil {
		return models.AllPost{}, err
	}
	return postData, nil

}

func (ju *jobseekerJobUseCase) UpdatePost(post models.EditPostReq) (models.EditPostRes, error) {

	okp, err := ju.postRepository.IsPostExistByPostId(post.PostId)

	if err != nil {
		return models.EditPostRes{}, err
	}

	if !okp {
		return models.EditPostRes{}, errors.New("post does not exist")
	}

	oku, err := ju.postRepository.IsPostExistByUserId(post.JobseekerId)

	if err != nil {
		return models.EditPostRes{}, err
	}

	if !oku {
		return models.EditPostRes{}, errors.New("post does not belongs to you")
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		return models.EditPostRes{}, err
	}

	h := helper.NewHelper(cfg)
	url, err := h.AddImageToAwsS3(post.Image)
	if err != nil {
		return models.EditPostRes{}, err
	}
	var editPost = models.EditPostRes{
		JobseekerId: post.JobseekerId,
		PostId:      uint(post.PostId),
		Title:       post.Title,
		Content:     post.Content,
		ImageUrl:    url,
	}
	PostData, err := ju.postRepository.UpdatePost(editPost)

	if err != nil {
		return models.EditPostRes{}, err
	}
	return PostData, nil

}

func (ju *jobseekerJobUseCase) DeletePost(postId, JobseekerId int) (bool, error) {

	okP, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		return false, err
	}
	if !okP {
		return false, errors.New("post does not exist")
	}

	okU, err := ju.postRepository.IsPostExistByUserId(JobseekerId)

	if err != nil {
		return false, err
	}
	if !okU {
		return false, errors.New("post is not belongs to you")
	}

	ok, err := ju.postRepository.DeletePost(postId, JobseekerId)

	if err != nil {
		return false, err
	}

	return ok, nil

}
func (ju *jobseekerJobUseCase) CreateCommentPost(postId, userId int, comment string) (bool, error) {

	if postId <= 0 {
		return false, errors.New("postId is required")
	}
	if comment == "" {
		return false, errors.New("comment is required")
	}

	okP, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		return false, err
	}
	if !okP {
		return false, errors.New("post does not exist")
	}
	ok, err := ju.postRepository.CreateCommentPost(postId, userId, comment)
	if err != nil {
		return false, err
	}
	return ok, nil

}

func (ju *jobseekerJobUseCase) UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error) {
	if postId <= 0 {
		return false, errors.New("postId is required")
	}
	if comment == "" {
		return false, errors.New("comment is required")
	}
	if commentId <= 0 {
		return false, errors.New("commentId is required")
	}

	okc, err := ju.postRepository.IsCommentIdExist(commentId)
	if err != nil {
		return false, err
	}
	if !okc {
		return false, errors.New("comment does not exist")
	}

	okcu, err := ju.postRepository.IsCommentIdBelongsUserId(commentId, userId)
	if err != nil {
		return false, err
	}
	if !okcu {
		return false, errors.New("comment does not belongs to you")
	}
	ok, err := ju.postRepository.UpdateCommentPost(commentId, postId, userId, comment)
	if err != nil {
		return false, err
	}
	return ok, nil

}

func (ju *jobseekerJobUseCase) DeleteCommentPost(postId, userId, commentId int) (bool, error) {
	if postId <= 0 {
		return false, errors.New("postId is required")
	}
	if commentId <= 0 {
		return false, errors.New("commentId is required")
	}
	okc, err := ju.postRepository.IsCommentIdExist(commentId)
	if err != nil {
		return false, err
	}
	if !okc {
		return false, errors.New("comment does not exist")
	}

	okcu, err := ju.postRepository.IsCommentIdBelongsUserId(commentId, userId)
	if err != nil {
		return false, err
	}
	if !okcu {
		return false, errors.New("comment does not belongs to you")
	}
	ok, err := ju.postRepository.DeleteCommentPost(commentId, postId, userId)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (ju *jobseekerJobUseCase) AddLikePost(postId, userId int) (bool, error) {
	okP, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		return false, err
	}
	if !okP {
		return false, errors.New("post does not exist")
	}
	okL, err := ju.postRepository.IsLikeExist(postId, userId)
	if err != nil {
		return false, err
	}
	if okL {
		return false, errors.New("you already like this post")
	}
	ok, err := ju.postRepository.AddLikePost(postId, userId)
	if err != nil {
		return false, err
	}
	return ok, nil
}

func (ju *jobseekerJobUseCase) RemoveLikePost(postId, userId int) (bool, error) {
	okP, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		return false, err
	}
	if !okP {
		return false, errors.New("post does not exist")
	}
	okL, err := ju.postRepository.IsLikeExist(postId, userId)
	if err != nil {
		return false, err
	}
	if !okL {
		return false, errors.New("you already unliked this post")
	}
	ok, err := ju.postRepository.RemoveLikePost(postId, userId)
	if err != nil {
		return false, err
	}
	return ok, nil
}