package client

import (
	"connectHub_gateway/pkg/client/post/interfaces"
	"connectHub_gateway/pkg/config"
	"connectHub_gateway/pkg/helper"
	jobseekerPb "connectHub_gateway/pkg/pb/post/jobseeker"
	"connectHub_gateway/pkg/utils/models"
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
)

type jobseekerPostClient struct {
	Client jobseekerPb.JobseekerPostServiceClient
}

func NewJobseekerPostClient(cfg config.Config) interfaces.JobseekerPostClient {

	grpcConnection, err := grpc.Dial(cfg.ConnetHubPost, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect to Post service", err)
	}

	grpcClient := jobseekerPb.NewJobseekerPostServiceClient(grpcConnection)

	return &jobseekerPostClient{Client: grpcClient}
}

func (jpc *jobseekerPostClient) CreatePost(post models.CreatePostReq) (models.CreatePostRes, error) {

	resp, err := jpc.Client.CreatePost(context.Background(), &jobseekerPb.CreatePostRequest{
		Title:    post.Title,
		Content:  post.Content,
		AuthorId: strconv.Itoa(post.JobseekerId),
		Image:    post.Image,
	})

	if err != nil {
		return models.CreatePostRes{}, err
	}

	jobseekerId, err := strconv.Atoi(resp.Post.AuthorId)

	if err != nil {
		return models.CreatePostRes{}, err
	}

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

	resp, err := jpc.Client.GetOnePost(context.Background(), &jobseekerPb.GetPostRequest{
		Id: uint64(postId),
	})
	if err != nil {
		return models.CreatePostRes{}, err
	}
	jobseekerId, err := strconv.Atoi(resp.Post.AuthorId)
	if err != nil {
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
	resp, err := jpc.Client.GetAllPost(context.Background(), &jobseekerPb.GetAllPostRequest{})
	if err != nil {
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
	return models.AllPost{Posts: posts}, nil
}

func (jpc *jobseekerPostClient) UpdatePost(post models.EditPostReq) (models.EditPostRes, error) {
	resp, err := jpc.Client.UpdatePost(context.Background(), &jobseekerPb.UpdatePostRequest{
		Id:          uint64(post.PostId),
		Title:       post.Title,
		Content:     post.Content,
		Image:       post.Image,
		JobseekerId: int64(post.JobseekerId),
	})
	if err != nil {
		return models.EditPostRes{}, err
	}
	jobseekerId, err := strconv.Atoi(resp.AuthorId)
	if err != nil {
		return models.EditPostRes{}, err
	}
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
	resp, err := jpc.Client.DeletePost(context.Background(), &jobseekerPb.DeletePostRequest{
		PostId:      uint64(postId),
		JobseekerId: uint64(JobseekerId),
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) CreateCommentPost(postId, userId int, comment string) (bool, error) {
	resp, err := jpc.Client.CreateCommentPost(context.Background(), &jobseekerPb.CreateCommentRequest{
		PostId:  uint64(postId),
		UserId:  uint64(userId),
		Comment: comment,
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error) {
	resp, err := jpc.Client.UpdateCommentPost(context.Background(), &jobseekerPb.UpdateCommentRequest{
		PostId:    uint64(postId),
		UserId:    uint64(userId),
		Comment:   comment,
		CommentId: uint64(commentId),
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) DeleteCommentPost(postId, userId, commentId int) (bool, error) {
	resp, err := jpc.Client.DeleteCommentPost(context.Background(), &jobseekerPb.DeleteCommentRequest{
		PostId:    uint64(postId),
		UserId:    uint64(userId),
		CommentId: uint64(commentId),
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) AddLikePost(postId, userId int) (bool, error) {
	resp, err := jpc.Client.AddLikePost(context.Background(), &jobseekerPb.AddLikeRequest{
		PostId: uint64(postId),
		UserId: uint64(userId),
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}

func (jpc *jobseekerPostClient) RemoveLikePost(postId, userId int) (bool, error) {
	resp, err := jpc.Client.RemoveLikePost(context.Background(), &jobseekerPb.RemoveLikeRequest{
		PostId: uint64(postId),
		UserId: uint64(userId),
	})
	if err != nil {
		return false, err
	}
	return resp.Success, nil
}
