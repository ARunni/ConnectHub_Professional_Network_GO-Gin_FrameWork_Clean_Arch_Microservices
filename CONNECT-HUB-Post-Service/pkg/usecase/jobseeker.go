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
	ju.Logger.Info("CreatePost at jobseekerJobUseCase started")

	var postReq models.CreatePostRes

	cfg, err := config.LoadConfig()
	if err != nil {
		ju.Logger.Error("error from LoadConfig", err)
		return models.CreatePostRes{}, err
	}

	h := helper.NewHelper(cfg)
	ju.Logger.Info("AddImageToAwsS3 started")
	url, err := h.AddImageToAwsS3(post.Image)
	if err != nil {
		ju.Logger.Error("error from AddImageToAwsS3", err)
		return models.CreatePostRes{}, err
	}
	ju.Logger.Info("AddImageToAwsS3 finished")
	postReq.Content = post.Content
	postReq.Title = post.Title
	postReq.JobseekerId = post.JobseekerId
	postReq.ImageUrl = url
	ju.Logger.Info("CreatePost at postRepository started")
	postData, err := ju.postRepository.CreatePost(postReq)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.CreatePostRes{}, err
	}
	ju.Logger.Info("CreatePost at postRepository finished")
	ju.Logger.Info("CreatePost at jobseekerJobUseCase finished")
	return postData, nil

}

func (ju *jobseekerJobUseCase) GetOnePost(postId int) (models.CreatePostRes, error) {
	ju.Logger.Info("GetOnePost at jobseekerJobUseCase started")
	ju.Logger.Info("IsPostExistByPostId at postRepository started")
	ok, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.CreatePostRes{}, err
	}
	ju.Logger.Info("IsPostExistByPostId at postRepository finished")
	if !ok {
		ju.Logger.Error("error : post does not exist")
		return models.CreatePostRes{}, errors.New("post does not exist")
	}
	ju.Logger.Info("GetOnePost at postRepository started")
	postData, err := ju.postRepository.GetOnePost(postId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.CreatePostRes{}, err
	}
	ju.Logger.Info("GetOnePost at postRepository finished")
	ju.Logger.Info("GetLikesCountPost at postRepository started")
	likeData, err := ju.postRepository.GetLikesCountPost(postId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.CreatePostRes{}, err
	}
	ju.Logger.Info("GetLikesCountPost at postRepository finished")
	ju.Logger.Info("GetCommentsPost at postRepository started")

	commentData, err := ju.postRepository.GetCommentsPost(postId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.CreatePostRes{}, err
	}
	ju.Logger.Info("GetCommentsPost at postRepository finished")
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
	ju.Logger.Info("GetOnePost at jobseekerJobUseCase finished")
	return result, nil
}

func (ju *jobseekerJobUseCase) GetAllPost() (models.AllPostData, error) {
	ju.Logger.Info("GetAllPost at jobseekerJobUseCase started")
	ju.Logger.Info("GetAllPost at postRepository started")
	postData, err := ju.postRepository.GetAllPost()
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.AllPostData{}, err
	}
	ju.Logger.Info("GetAllPost at postRepository finished")
	var postAllData []models.CreatePostRes

	for _, post := range postData.Posts {
		ju.Logger.Info("GetLikesCountPost at postRepository started")
		likeData, err := ju.postRepository.GetLikesCountPost(post.ID)
		if err != nil {
			ju.Logger.Error("error from postRepository", err)
			return models.AllPostData{}, err
		}
		ju.Logger.Info("GetLikesCountPost at postRepository finished")
		ju.Logger.Info("GetCommentsPost at postRepository started")
		commentData, err := ju.postRepository.GetCommentsPost(post.ID)
		if err != nil {
			ju.Logger.Error("error from postRepository", err)
			return models.AllPostData{}, err
		}
		ju.Logger.Info("GetCommentsPost at postRepository finished")
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
	ju.Logger.Info("GetAllPost at jobseekerJobUseCase finished")
	return models.AllPostData{Posts: postAllData}, nil
}

//

func (ju *jobseekerJobUseCase) UpdatePost(post models.EditPostReq) (models.EditPostRes, error) {
	ju.Logger.Info("UpdatePost at jobseekerJobUseCase started")
	ju.Logger.Info("IsPostExistByPostId at postRepository started")
	okp, err := ju.postRepository.IsPostExistByPostId(post.PostId)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.EditPostRes{}, err
	}
	ju.Logger.Info("IsPostExistByPostId at postRepository finished")
	if !okp {
		ju.Logger.Error("error : post does not exist")
		return models.EditPostRes{}, errors.New("post does not exist")
	}
	ju.Logger.Info("IsPostExistByUserId at postRepository started")
	oku, err := ju.postRepository.IsPostExistByUserId(post.JobseekerId)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.EditPostRes{}, err
	}
	ju.Logger.Info("IsPostExistByUserId at postRepository finished")
	if !oku {
		ju.Logger.Error("error : post does not belongs to you")
		return models.EditPostRes{}, errors.New("post does not belongs to you")
	}
	ju.Logger.Info("LoadConfig at config started")
	cfg, err := config.LoadConfig()
	if err != nil {
		ju.Logger.Error("error from LoadConfig", err)
		return models.EditPostRes{}, err
	}
	ju.Logger.Info("LoadConfig at config finished")
	ju.Logger.Info("NewHelper at helper started")
	h := helper.NewHelper(cfg)
	ju.Logger.Info("NewHelper at helper finished")
	ju.Logger.Info("AddImageToAwsS3 started")
	url, err := h.AddImageToAwsS3(post.Image)
	if err != nil {
		ju.Logger.Error("error from AddImageToAwsS3", err)
		return models.EditPostRes{}, err
	}
	ju.Logger.Info("AddImageToAwsS3 finished")
	var editPost = models.EditPostRes{
		JobseekerId: post.JobseekerId,
		PostId:      uint(post.PostId),
		Title:       post.Title,
		Content:     post.Content,
		ImageUrl:    url,
	}
	ju.Logger.Info("UpdatePost at postRepository started")
	PostData, err := ju.postRepository.UpdatePost(editPost)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return models.EditPostRes{}, err
	}
	ju.Logger.Info("UpdatePost at postRepository finished")
	ju.Logger.Info("UpdatePost at jobseekerJobUseCase finished")
	return PostData, nil

}

func (ju *jobseekerJobUseCase) DeletePost(postId, JobseekerId int) (bool, error) {
	ju.Logger.Info("DeletePost at jobseekerJobUseCase started")
	ju.Logger.Info("IsPostExistByPostId at postRepository started")
	okP, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("IsPostExistByPostId at postRepository finished")
	if !okP {
		ju.Logger.Error("error : post does not exist")
		return false, errors.New("post does not exist")
	}
	ju.Logger.Info("IsPostExistByUserId at postRepository started")
	okU, err := ju.postRepository.IsPostExistByUserId(JobseekerId)

	if err != nil {
		ju.Logger.Error("error from postRepository")
		return false, err
	}
	ju.Logger.Info("IsPostExistByUserId at postRepository finished")
	if !okU {
		ju.Logger.Error("error : post is not belongs to you")
		return false, errors.New("post is not belongs to you")
	}
	ju.Logger.Info("DeletePost at postRepository started")
	ok, err := ju.postRepository.DeletePost(postId, JobseekerId)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("DeletePost at postRepository finished")
	ju.Logger.Info("DeletePost at jobseekerJobUseCase finished")
	return ok, nil

}
func (ju *jobseekerJobUseCase) CreateCommentPost(postId, userId int, comment string) (bool, error) {
	ju.Logger.Info("CreateCommentPost at jobseekerJobUseCase started")
	if postId <= 0 {
		ju.Logger.Error("error : postId is required")
		return false, errors.New("postId is required")
	}
	if comment == "" {
		ju.Logger.Error("error : comment is required")
		return false, errors.New("comment is required")
	}
	ju.Logger.Info("IsPostExistByPostId at postRepository started")
	okP, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("IsPostExistByPostId at postRepository finished")
	if !okP {
		ju.Logger.Error("error : post does not exist")
		return false, errors.New("post does not exist")
	}
	ju.Logger.Info("CreateCommentPost at postRepository started")
	ok, err := ju.postRepository.CreateCommentPost(postId, userId, comment)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("CreateCommentPost at postRepository finished")
	//
	ju.Logger.Info("UserData at authClient started")
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
	ju.Logger.Info("UpdateCommentPost at jobseekerJobUseCase started")
	if postId <= 0 {
		ju.Logger.Error("error : postId is required")
		return false, errors.New("postId is required")
	}
	if comment == "" {
		ju.Logger.Error("error : comment is required")
		return false, errors.New("comment is required")
	}
	if commentId <= 0 {
		ju.Logger.Error("error : commentId is required")
		return false, errors.New("commentId is required")
	}
	ju.Logger.Info("IsCommentIdExist at postRepository started")
	okc, err := ju.postRepository.IsCommentIdExist(commentId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("IsCommentIdExist at postRepository finished")
	if !okc {
		ju.Logger.Error("error : comment does not exist")
		return false, errors.New("comment does not exist")
	}
	ju.Logger.Info("IsCommentIdBelongsUserId at postRepository started")
	okcu, err := ju.postRepository.IsCommentIdBelongsUserId(commentId, userId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("IsCommentIdBelongsUserId at postRepository finished")
	if !okcu {
		ju.Logger.Error("error : comment does not belongs to you")
		return false, errors.New("comment does not belongs to you")
	}
	ju.Logger.Info("UpdateCommentPost at postRepository started")
	ok, err := ju.postRepository.UpdateCommentPost(commentId, postId, userId, comment)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("UpdateCommentPost at postRepository finished")
	ju.Logger.Info("UserData at authClient started")
	senderData, err := ju.authClient.UserData(userId)
	if err != nil {
		ju.Logger.Error("error from authClient", err)
		return false, err
	}
	ju.Logger.Info("UserData at authClient finished")
	ju.Logger.Info("GetOnePost at postRepository started")
	postUser, err := ju.postRepository.GetOnePost(postId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("GetOnePost at postRepository finished")
	msg := fmt.Sprintf("%s comment edited on Your PostID %d", senderData.Username, postId)
	ju.Logger.Info("SendNotification at helper started")
	helper.SendNotification(models.Notification{
		UserID:     postUser.JobseekerId,
		SenderID:   senderData.UserId,
		PostID:     postId,
		SenderName: senderData.Username,
	}, []byte(msg))
	ju.Logger.Info("SendNotification at helper finished")
	ju.Logger.Info("UpdateCommentPost at jobseekerJobUseCase finished")
	return ok, nil

}

func (ju *jobseekerJobUseCase) DeleteCommentPost(postId, userId, commentId int) (bool, error) {
	ju.Logger.Info("DeleteCommentPost at jobseekerJobUseCase started")
	if postId <= 0 {
		ju.Logger.Error("error : postId is required")
		return false, errors.New("postId is required")
	}
	if commentId <= 0 {
		ju.Logger.Error("error : commentId is required")
		return false, errors.New("commentId is required")
	}
	ju.Logger.Info("IsCommentIdExist at postRepository started")
	okc, err := ju.postRepository.IsCommentIdExist(commentId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("IsCommentIdExist at postRepository finished")
	if !okc {
		ju.Logger.Error("error : comment does not exist")
		return false, errors.New("comment does not exist")
	}
	ju.Logger.Info("IsCommentIdBelongsUserId at postRepository started")
	okcu, err := ju.postRepository.IsCommentIdBelongsUserId(commentId, userId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("IsCommentIdBelongsUserId at postRepository finished")
	if !okcu {
		ju.Logger.Error("error : comment does not belongs to you")
		return false, errors.New("comment does not belongs to you")
	}
	ju.Logger.Info("DeleteCommentPost at postRepository started")
	ok, err := ju.postRepository.DeleteCommentPost(commentId, postId, userId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("DeleteCommentPost at postRepository finished")
	ju.Logger.Info("DeleteCommentPost at jobseekerJobUseCase finished")
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
	ju.Logger.Info("RemoveLikePost at jobseekerJobUseCase started")
	ju.Logger.Info("IsPostExistByPostId at postRepository started")
	okP, err := ju.postRepository.IsPostExistByPostId(postId)

	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("IsPostExistByPostId at postRepository finished")
	if !okP {
		ju.Logger.Error("error : post does not exist")
		return false, errors.New("post does not exist")
	}
	ju.Logger.Info("IsLikeExist at postRepository started")
	okL, err := ju.postRepository.IsLikeExist(postId, userId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("IsLikeExist at postRepository finished")
	if !okL {
		ju.Logger.Error("error : you already unliked this post")
		return false, errors.New("you already unliked this post")
	}
	ju.Logger.Info("RemoveLikePost at postRepository started")
	ok, err := ju.postRepository.RemoveLikePost(postId, userId)
	if err != nil {
		ju.Logger.Error("error from postRepository", err)
		return false, err
	}
	ju.Logger.Info("RemoveLikePost at postRepository finished")
	ju.Logger.Info("RemoveLikePost at jobseekerJobUseCase finished")
	return ok, nil
}
