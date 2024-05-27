package service

import (
	jobseekerPb "ConnetHub_post/pkg/pb/post/jobseeker"
	"ConnetHub_post/pkg/usecase/interfaces"
	"ConnetHub_post/pkg/utils/models"
	"context"
	"strconv"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type JobseekerPostServer struct {
	postUseCase interfaces.JobseekerPostUsecase
	jobseekerPb.UnimplementedJobseekerPostServiceServer
}

func NewJobseekerPostServer(useCase interfaces.JobseekerPostUsecase) jobseekerPb.JobseekerPostServiceServer {

	return &JobseekerPostServer{
		postUseCase: useCase,
	}
}

func (jps *JobseekerPostServer) CreatePost(ctx context.Context, req *jobseekerPb.CreatePostRequest) (*jobseekerPb.CreatePostResponse, error) {

	userId, err := strconv.Atoi(req.AuthorId)
	if err != nil {
		return nil, err
	}
	var post = models.CreatePostReq{
		JobseekerId: userId,
		Title:       req.Title,
		Content:     req.Content,
		Image:       req.Image,
	}
	postData, err := jps.postUseCase.CreatePost(post)
	if err != nil {
		return nil, err
	}
	user := strconv.Itoa(postData.JobseekerId)
	return &jobseekerPb.CreatePostResponse{
		Post: &jobseekerPb.Post{
			Id:        uint64(postData.ID),
			Title:     postData.Title,
			Content:   postData.Content,
			AuthorId:  user,
			CreatedAt: timestamppb.New(postData.CreatedAt),
			Url:       postData.ImageUrl,
		},
	}, nil
}

func (jps *JobseekerPostServer) GetAllPost(ctx context.Context, req *jobseekerPb.GetAllPostRequest) (*jobseekerPb.GetAllPostResponse, error) {

	postData, err := jps.postUseCase.GetAllPost()
	if err != nil {
		return nil, err
	}
	var posts []*jobseekerPb.Post
	for _, post := range postData.Posts {
		posts = append(posts, &jobseekerPb.Post{
			Id:        uint64(post.ID),
			Title:     post.Title,
			Content:   post.Content,
			AuthorId:  strconv.Itoa(post.JobseekerId),
			CreatedAt: timestamppb.New(post.CreatedAt),
			Url:       post.ImageUrl,
		})
	}
	return &jobseekerPb.GetAllPostResponse{
		Posts: posts,
	}, nil

}

func (jps *JobseekerPostServer) GetPost(ctx context.Context, req *jobseekerPb.GetPostRequest) (*jobseekerPb.GetPostResponse, error) {

	postData, err := jps.postUseCase.GetOnePost(int(req.Id))
	if err != nil {
		return nil, err
	}

	var post = &jobseekerPb.GetPostResponse{
		Post: &jobseekerPb.Post{
			Id:        uint64(postData.ID),
			Title:     postData.Title,
			Content:   postData.Content,
			AuthorId:  strconv.Itoa(postData.JobseekerId),
			CreatedAt: timestamppb.New(postData.CreatedAt),
			Url:       postData.ImageUrl,
		},
	}
	return post, nil

}

func (jps *JobseekerPostServer) UpdatePost(ctx context.Context, req *jobseekerPb.UpdatePostRequest) (*jobseekerPb.UpdatePostResponse, error) {
	var editPost = models.EditPostReq{
		JobseekerId: int(req.JobseekerId),
		Title:       req.Title,
		PostId:      int(req.Id),
		Content:     req.Content,
		Image:       req.Image,
	}
	postData, err := jps.postUseCase.UpdatePost(editPost)
	if err != nil {
		return nil, err
	}

	return &jobseekerPb.UpdatePostResponse{

		Id:        uint64(postData.PostId),
		Title:     postData.Title,
		Content:   postData.Content,
		AuthorId:  strconv.Itoa(postData.JobseekerId),
		UpdatedAt: timestamppb.New(postData.UpdatedAt),
		Url:       postData.ImageUrl,
	}, nil

}

func (jps *JobseekerPostServer) DeletePost(ctx context.Context, req *jobseekerPb.DeletePostRequest) (*jobseekerPb.DeletePostResponse, error) {

	postData, err := jps.postUseCase.DeletePost(int(req.PostId), int(req.JobseekerId))
	if err != nil {
		return nil, err
	}

	return &jobseekerPb.DeletePostResponse{
		Success: postData,
	}, nil

}
