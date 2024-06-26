package client

import (
	logging "github.com/ARunni/connectHub_gateway/Logging"
	"github.com/ARunni/connectHub_gateway/pkg/client/post/interfaces"
	"github.com/ARunni/connectHub_gateway/pkg/config"
	"github.com/ARunni/connectHub_gateway/pkg/helper"
	jobseekerPb "github.com/ARunni/connectHub_gateway/pkg/pb/post/jobseeker"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type jobseekerPostClient struct {
	Client  jobseekerPb.JobseekerPostServiceClient
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewJobseekerPostClient(cfg config.Config) interfaces.JobseekerPostClient {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	grpcConnection, err := grpc.Dial(cfg.ConnetHubPost, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect to Post service", err)
	}

	grpcClient := jobseekerPb.NewJobseekerPostServiceClient(grpcConnection)

	return &jobseekerPostClient{
		Client:  grpcClient,
		Logger:  logger,
		LogFile: logFile,
	}
}

func (jpc *jobseekerPostClient) CreatePost(post models.CreatePostReq) (models.CreatePostRes, error) {

	jpc.Logger.Info("CreatePost at client started")

	resp, err := jpc.Client.CreatePost(context.Background(), &jobseekerPb.CreatePostRequest{
		Title:    post.Title,
		Content:  post.Content,
		AuthorId: strconv.Itoa(post.JobseekerId),
		Image:    post.Image,
	})

	if err != nil {
		jpc.Logger.Error("Error creating post: ", err)
		return models.CreatePostRes{}, err
	}

	jobseekerId, err := strconv.Atoi(resp.Post.AuthorId)

	if err != nil {
		jpc.Logger.Error("Error converting authorId to int: ", err)
		return models.CreatePostRes{}, err
	}
	jpc.Logger.Info("CreatePost at client success")

	return models.CreatePostRes{
		ID:          int(resp.Post.Id),
		JobseekerId: jobseekerId,
		Title:       resp.Post.Title,
		Content:     resp.Post.Content,
		ImageUrl:    resp.Post.Url,
		CreatedAt:   helper.FromProtoTimestamp(resp.Post.CreatedAt),
		Comments:    nil,
		Likes:       int(resp.Post.Likes),
	}, nil
}

func (jpc *jobseekerPostClient) GetOnePost(postId int) (models.CreatePostRes, error) {

	jpc.Logger.Info("GetOnePost at client started")

	resp, err := jpc.Client.GetOnePost(context.Background(), &jobseekerPb.GetPostRequest{
		Id: uint64(postId),
	})
	if err != nil {
		jpc.Logger.Error("Error getting post: ", err)
		return models.CreatePostRes{}, err
	}
	jobseekerId, err := strconv.Atoi(resp.Post.AuthorId)
	if err != nil {
		jpc.Logger.Error("Error converting authorId to int: ", err)
		return models.CreatePostRes{}, err
	}
	var commentDatas []models.CommentData
	for _, commentD := range resp.Post.Comments {
		commentDatas = append(commentDatas, models.CommentData{
			ID:          uint(commentD.Id),
			Comment:     commentD.Comment,
			JobseekerId: uint(commentD.AuthorId),
			CreatedAt:   helper.FromProtoTimestamp(commentD.CreatedAt),
			UpdatedAt:   helper.FromProtoTimestamp(commentD.UpdatedAt),
		})
	}
	jpc.Logger.Info("GetOnePost at client success")
	return models.CreatePostRes{
		ID:          int(resp.Post.Id),
		JobseekerId: jobseekerId,
		Title:       resp.Post.Title,
		Content:     resp.Post.Content,
		ImageUrl:    resp.Post.Url,
		CreatedAt:   helper.FromProtoTimestamp(resp.Post.CreatedAt),
		Comments:    commentDatas,
		Likes:       int(resp.Post.Likes),
	}, nil
}

func (jpc *jobseekerPostClient) GetAllPost() (models.AllPost, error) {
	jpc.Logger.Info("GetAllPost at client started")
	resp, err := jpc.Client.GetAllPost(context.Background(), &jobseekerPb.GetAllPostRequest{})
	if err != nil {
		jpc.Logger.Error("Error getting all posts: ", err)
		return models.AllPost{}, err
	}
	var posts []models.CreatePostRes
	var commentData []models.CommentData
	for _, post := range resp.Posts {
		for _, Comment := range post.Comments {
			commentData = append(commentData, models.CommentData{
				ID:          uint(Comment.Id),
				Comment:     Comment.Comment,
				JobseekerId: uint(Comment.AuthorId),
				CreatedAt:   helper.FromProtoTimestamp(Comment.CreatedAt),
				UpdatedAt:   helper.FromProtoTimestamp(Comment.UpdatedAt),
			})
		}
		createdAt := helper.FromProtoTimestamp(post.CreatedAt)
		jobseekerId, err := strconv.Atoi(post.AuthorId)
		if err != nil {
			jpc.Logger.Error("Error converting authorId to int: ", err)
			return models.AllPost{}, err
		}
		posts = append(posts, models.CreatePostRes{
			ID:          int(post.Id),
			JobseekerId: jobseekerId,
			Title:       post.Title,
			Content:     post.Content,
			ImageUrl:    post.Url,
			CreatedAt:   createdAt,
			Comments:    commentData,
			Likes:       int(post.Likes),
		})
	}
	jpc.Logger.Info("GetAllPost at client success")
	return models.AllPost{Posts: posts}, nil
}

func (jpc *jobseekerPostClient) UpdatePost(post models.EditPostReq) (models.EditPostRes, error) {
	jpc.Logger.Info("UpdatePost at client started")
	resp, err := jpc.Client.UpdatePost(context.Background(), &jobseekerPb.UpdatePostRequest{
		Id:          uint64(post.PostId),
		Title:       post.Title,
		Content:     post.Content,
		Image:       post.Image,
		JobseekerId: int64(post.JobseekerId),
	})
	if err != nil {
		jpc.Logger.Error("Error updating post: ", err)
		return models.EditPostRes{}, err
	}
	jobseekerId, err := strconv.Atoi(resp.AuthorId)
	if err != nil {
		jpc.Logger.Error("Error converting authorId to int: ", err)
		return models.EditPostRes{}, err
	}
	jpc.Logger.Info("UpdatePost at client success")
	return models.EditPostRes{
		JobseekerId: jobseekerId,
		PostId:      int(resp.Id),
		Title:       resp.Title,
		Content:     resp.Content,
		ImageUrl:    resp.Url,
		UpdatedAt:   helper.FromProtoTimestamp(resp.UpdatedAt),
	}, nil
}

func (jpc *jobseekerPostClient) DeletePost(postId, JobseekerId int) (bool, error) {
	jpc.Logger.Info("DeletePost at client started")
	resp, err := jpc.Client.DeletePost(context.Background(), &jobseekerPb.DeletePostRequest{
		PostId:      uint64(postId),
		JobseekerId: uint64(JobseekerId),
	})
	if err != nil {
		jpc.Logger.Error("Error deleting post: ", err)
		return false, err
	}
	jpc.Logger.Info("DeletePost at client success")
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) CreateCommentPost(postId, userId int, comment string) (bool, error) {
	jpc.Logger.Info("CreateCommentPost at client started")
	resp, err := jpc.Client.CreateCommentPost(context.Background(), &jobseekerPb.CreateCommentRequest{
		PostId:  uint64(postId),
		UserId:  uint64(userId),
		Comment: comment,
	})
	if err != nil {
		jpc.Logger.Error("Error creating comment: ", err)
		return false, err
	}
	jpc.Logger.Info("CreateCommentPost at client success")
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error) {
	jpc.Logger.Info("UpdateCommentPost at client started")
	resp, err := jpc.Client.UpdateCommentPost(context.Background(), &jobseekerPb.UpdateCommentRequest{
		PostId:    uint64(postId),
		UserId:    uint64(userId),
		Comment:   comment,
		CommentId: uint64(commentId),
	})
	if err != nil {
		jpc.Logger.Error("Error updating comment: ", err)
		return false, err
	}
	jpc.Logger.Info("UpdateCommentPost at client success")
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) DeleteCommentPost(postId, userId, commentId int) (bool, error) {
	jpc.Logger.Info("UpdateCommentPost at client started")
	resp, err := jpc.Client.DeleteCommentPost(context.Background(), &jobseekerPb.DeleteCommentRequest{
		PostId:    uint64(postId),
		UserId:    uint64(userId),
		CommentId: uint64(commentId),
	})
	if err != nil {
		jpc.Logger.Error("Error deleting comment: ", err)
		return false, err
	}
	jpc.Logger.Info("UpdateCommentPost at client success")
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) AddLikePost(postId, userId int) (bool, error) {
	jpc.Logger.Info("AddLikePost at client started")
	resp, err := jpc.Client.AddLikePost(context.Background(), &jobseekerPb.AddLikeRequest{
		PostId: uint64(postId),
		UserId: uint64(userId),
	})
	if err != nil {
		jpc.Logger.Error("Error adding like: ", err)
		return false, err
	}
	jpc.Logger.Info("AddLikePost at client success")
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) RemoveLikePost(postId, userId int) (bool, error) {
	jpc.Logger.Info("RemoveLikePost at client started")
	resp, err := jpc.Client.RemoveLikePost(context.Background(), &jobseekerPb.RemoveLikeRequest{
		PostId: uint64(postId),
		UserId: uint64(userId),
	})
	if err != nil {
		jpc.Logger.Error("Error removing like: ", err)
		return false, err
	}
	jpc.Logger.Info("RemoveLikePost at client success")
	return resp.Success, nil
}
