package usecase

import (
	"errors"
	"fmt"
	"os"

	logging "github.com/ARunni/ConnetHub_post/Logging"
	auth "github.com/ARunni/ConnetHub_post/pkg/client/auth/interfaces"
	"github.com/ARunni/ConnetHub_post/pkg/config"
	"github.com/ARunni/ConnetHub_post/pkg/helper"
	repo "github.com/ARunni/ConnetHub_post/pkg/repository/interfaces"
	usecase "github.com/ARunni/ConnetHub_post/pkg/usecase/interfaces"
	"github.com/ARunni/ConnetHub_post/pkg/utils/models"

	"github.com/sirupsen/logrus"
)

type jobseekerJobUseCase struct {
	postRepository repo.JobseekerPostRepository
	authClient     auth.Newauthclient
	Logger         *logrus.Logger
	LogFile        *os.File
}

func NewJobseekerpostUseCase(repo repo.JobseekerPostRepository, client auth.Newauthclient) usecase.JobseekerPostUsecase {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Post.log")
	return &jobseekerJobUseCase{
		postRepository: repo,
		Logger:         logger,
		LogFile:        logFile,
		authClient:     client,
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
	likeData, err := ju.postRepository.GetLikesCountPost(postId)
	if err != nil {
		return models.CreatePostRes{}, err
	}
	commentData, err := ju.postRepository.GetCommentsPost(postId)
	if err != nil {
		return models.CreatePostRes{}, err
	}
	var result = models.CreatePostRes{
		ID:          postData.ID,
		JobseekerId: postData.JobseekerId,
		Title:       postData.Title,
		Content:     postData.Content,
		ImageUrl:    postData.ImageUrl,
		Comments:    commentData,
		Likes:       likeData,
		CreatedAt:   postData.CreatedAt,
	}
	// postData.Likes = likeData
	// postData.Comments = commentData

	return result, nil
}

func (ju *jobseekerJobUseCase) GetAllPost() (models.AllPostData, error) {

	postData, err := ju.postRepository.GetAllPost()
	if err != nil {
		return models.AllPostData{}, err
	}

	var postAllData []models.CreatePostRes

	for _, post := range postData.Posts {
		likeData, err := ju.postRepository.GetLikesCountPost(post.ID)
		if err != nil {
			return models.AllPostData{}, err
		}
		commentData, err := ju.postRepository.GetCommentsPost(post.ID)
		if err != nil {
			return models.AllPostData{}, err
		}

		postAllData = append(postAllData, models.CreatePostRes{
			ID:          post.ID,
			JobseekerId: post.JobseekerId,
			Title:       post.Title,
			Content:     post.Content,
			ImageUrl:    post.ImageUrl,
			Comments:    commentData,
			Likes:       likeData,
			CreatedAt:   post.CreatedAt,
		})
	}

	return models.AllPostData{Posts: postAllData}, nil
}

//

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

	//

	senderData, err := ju.authClient.UserData(userId)
	if err != nil {
		ju.Logger.Error("error from authClient", err)
		return false, err
	}
	postUser, err := ju.postRepository.GetOnePost(postId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	msg := fmt.Sprintf("%s comented on Your PostID %d", senderData.Username, postId)
	helper.SendNotification(models.Notification{
		UserID:     postUser.JobseekerId,
		SenderID:   senderData.UserId,
		PostID:     postId,
		SenderName: senderData.Username,
	}, []byte(msg))

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

	senderData, err := ju.authClient.UserData(userId)
	if err != nil {
		ju.Logger.Error("error from authClient", err)
		return false, err
	}
	postUser, err := ju.postRepository.GetOnePost(postId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	msg := fmt.Sprintf("%s comment edited on Your PostID %d", senderData.Username, postId)
	helper.SendNotification(models.Notification{
		UserID:     postUser.JobseekerId,
		SenderID:   senderData.UserId,
		PostID:     postId,
		SenderName: senderData.Username,
	}, []byte(msg))

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
	ju.Logger.Info("AddLikePost at jobseekerJobUseCase started")
	if postId <= 0 {
		ju.Logger.Error("error postId is required")
		return false, errors.New("postId is required")
	}
	okP, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	if !okP {
		ju.Logger.Error("error postId does not exist")
		return false, errors.New("post does not exist")
	}
	okL, err := ju.postRepository.IsLikeExist(postId, userId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	if okL {
		ju.Logger.Error("error you already like this post")
		return false, errors.New("you already like this post")
	}
	senderData, err := ju.authClient.UserData(userId)
	if err != nil {
		ju.Logger.Error("error from authClient", err)
		return false, err
	}

	ok, err := ju.postRepository.AddLikePost(postId, userId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	postUser, err := ju.postRepository.GetOnePost(postId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	msg := fmt.Sprintf("%s Liked on Your PostID %d", senderData.Username, postId)
	helper.SendNotification(models.Notification{
		UserID:     postUser.JobseekerId,
		SenderID:   senderData.UserId,
		PostID:     postId,
		SenderName: senderData.Username,
	}, []byte(msg))
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
