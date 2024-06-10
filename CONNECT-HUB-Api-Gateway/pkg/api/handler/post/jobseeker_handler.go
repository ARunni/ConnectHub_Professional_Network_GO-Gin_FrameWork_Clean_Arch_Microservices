package handler

import (
	"connectHub_gateway/pkg/client/post/interfaces"
	"connectHub_gateway/pkg/utils/models"
	"connectHub_gateway/pkg/utils/response"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	msg "github.com/ARunni/Error_Message"

	"github.com/gin-gonic/gin"
)

type JobseekerPostHandler struct {
	GRPC_Client interfaces.JobseekerPostClient
}

func NewJobseekerPostHandler(grpc_client interfaces.JobseekerPostClient) *JobseekerPostHandler {
	return &JobseekerPostHandler{
		GRPC_Client: grpc_client,
	}
}

// CreatePost creates a new post by a jobseeker.
// @Summary Create a post
// @Description Create a new post by a jobseeker
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param title formData string true "Title of the post"
// @Param content formData string true "Content of the post"
// @Param image formData file true "Image file for the post"
// @Success 200 {object} response.Response "Post created successfully"
// @Failure 400 {object} response.Response "Failed to create post: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to create post"
// @Router /jobseeker/post [post]
func (jph *JobseekerPostHandler) CreatePost(c *gin.Context) {

	var post models.CreatePostReq

	post.Title = c.PostForm("title")
	post.Content = c.PostForm("content")
	jobseekerIdAny, ok := c.Get("id")
	if !ok {
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	jobseekerId := jobseekerIdAny.(int)
	post.JobseekerId = jobseekerId

	file, err := c.FormFile("image")
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Getting Image Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	fileContent, err := file.Open()
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Error opening the file", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return

	}
	defer fileContent.Close()
	imageData, err := ioutil.ReadAll(fileContent)
	if err != nil {

		errResp := response.ClientResponse(http.StatusInternalServerError, "Error reading the file", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	post.Image = imageData

	postData, err := jph.GRPC_Client.CreatePost(post)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgAddSuccess, postData, nil)
	c.JSON(http.StatusOK, successRes)

}

// GetOnePost retrieves a single post by its ID.
// @Summary Get a post by ID
// @Description Retrieve a single post by its ID
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param post_id query int true "ID of the post to retrieve"
// @Success 200 {object} response.Response "Post retrieved successfully"
// @Failure 400 {object} response.Response "Failed to get post: missing or incorrect parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to get post"
// @Router /jobseeker/post [get]
func (jph *JobseekerPostHandler) GetOnePost(c *gin.Context) {
	postIdstr := c.Query("post_id")
	postId, err := strconv.Atoi(postIdstr)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrDatatypeConversion, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	postData, err := jph.GRPC_Client.GetOnePost(postId)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, postData, nil)
	c.JSON(http.StatusOK, successRes)

}

// GetAllPost retrieves all posts.
// @Summary Get all posts
// @Description Retrieve all posts
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} response.Response "Posts retrieved successfully"
// @Failure 500 {object} response.Response "Internal server error: failed to get posts"
// @Router /jobseeker/posts [get]
func (jph *JobseekerPostHandler) GetAllPost(c *gin.Context) {

	postData, err := jph.GRPC_Client.GetAllPost()
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgGetSucces, postData, nil)
	c.JSON(http.StatusOK, successRes)

}

// UpdatePost updates a post.
// @Summary Update a post
// @Description Update a post by providing post ID, title, content, and image
// @Tags Jobseeker
// @Accept multipart/form-data
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param post_id formData string true "Post ID"
// @Param title formData string true "Post title"
// @Param content formData string true "Post content"
// @Param image formData file true "Post image"
// @Success 200 {object} response.Response "Post updated successfully"
// @Failure 400 {object} response.Response "Bad request: failed to convert post ID"
// @Failure 500 {object} response.Response "Internal server error: failed to update post"
// @Router /jobseeker/posts/update [post]
func (jph *JobseekerPostHandler) UpdatePost(c *gin.Context) {
	var post models.EditPostReq

	postIdstr := c.PostForm("post_id")
	// fmt.Println("ajkhgflkgsjha", postIdstr)
	postId, err := strconv.Atoi(postIdstr)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.MsgConvErr, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	post.PostId = postId
	post.Title = c.PostForm("title")
	post.Content = c.PostForm("content")
	jobseekerIdAny, ok := c.Get("id")
	if !ok {
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	jobseekerId := jobseekerIdAny.(int)
	post.JobseekerId = jobseekerId

	file, err := c.FormFile("image")
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Getting Image Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	fileContent, err := file.Open()
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, "Error opening the file", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return

	}
	defer fileContent.Close()
	imageData, err := ioutil.ReadAll(fileContent)
	if err != nil {

		errResp := response.ClientResponse(http.StatusInternalServerError, "Error reading the file", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	post.Image = imageData
	postData, err := jph.GRPC_Client.UpdatePost(post)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgUpdateSuccess, postData, nil)
	c.JSON(http.StatusOK, successRes)

}

// DeletePost deletes a post.
// @Summary Delete a post
// @Description Delete a post by providing post ID
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param post_id query string true "Post ID"
// @Success 200 {object} response.Response "Post deleted successfully"
// @Failure 400 {object} response.Response "Bad request: failed to convert post ID"
// @Failure 500 {object} response.Response "Internal server error: failed to delete post"
// @Router /jobseeker/posts/delete [delete]
func (jph *JobseekerPostHandler) DeletePost(c *gin.Context) {
	postIdstr := c.Query("post_id")
	postId, err := strconv.Atoi(postIdstr)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrDatatypeConversion, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	jobseekerIdAny, ok := c.Get("id")
	if !ok {
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	jobseekerId := jobseekerIdAny.(int)
	postData, err := jph.GRPC_Client.DeletePost(postId, jobseekerId)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, postData, nil)
	c.JSON(http.StatusOK, successRes)

}

// CreateCommentPost creates a comment on a post.
// @Summary Create a comment on a post
// @Description Create a comment on a post by providing post ID and comment content
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param request body models.CreateCommentPost true "Comment request body"
// @Success 200 {object} response.Response "Comment created successfully"
// @Failure 400 {object} response.Response "Bad request: incorrect format"
// @Failure 500 {object} response.Response "Internal server error: failed to create comment"
// @Router /jobseeker/posts/comment [post]
func (jph *JobseekerPostHandler) CreateCommentPost(c *gin.Context) {
	userIdany, ok := c.Get("id")
	if !ok {
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	var comment models.CreateCommentPost
	userId := userIdany.(int)
	comment.UserId = userId
	if err := c.ShouldBindJSON(&comment); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	postOk, err := jph.GRPC_Client.CreateCommentPost(comment.PostId, comment.UserId, comment.Comment)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, postOk, nil)
	c.JSON(http.StatusOK, successRes)

}

// UpdateCommentPost updates a comment on a post.
// @Summary Update a comment on a post
// @Description Update a comment on a post by providing comment ID, post ID, and updated comment content
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param comment_id path int true "Comment ID"
// @Param post_id path int true "Post ID"
// @Param request body models.UpdateCommentPost true "Comment request body"
// @Success 200 {object} response.Response "Comment updated successfully"
// @Failure 400 {object} response.Response "Bad request: incorrect format"
// @Failure 500 {object} response.Response "Internal server error: failed to update comment"
// @Router /jobseeker/posts/comment/update [put]
func (jph *JobseekerPostHandler) UpdateCommentPost(c *gin.Context) {
	userIdany, ok := c.Get("id")
	if !ok {
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	var comment models.UpdateCommentPost
	userId := userIdany.(int)
	comment.UserId = userId
	if err := c.ShouldBindJSON(&comment); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	postOk, err := jph.GRPC_Client.UpdateCommentPost(comment.CommentId, comment.PostId, comment.UserId, comment.Comment)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, postOk, nil)
	c.JSON(http.StatusOK, successRes)

}

// DeleteCommentPost deletes a comment on a post.
// @Summary Delete a comment on a post
// @Description Delete a comment on a post by providing post ID and comment ID
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param post_id path int true "Post ID"
// @Param comment_id path int true "Comment ID"
// @Success 200 {object} response.Response "Comment deleted successfully"
// @Failure 400 {object} response.Response "Bad request: incorrect format"
// @Failure 500 {object} response.Response "Internal server error: failed to delete comment"
// @Router /jobseeker/posts/comment/delete [delete]
func (jph *JobseekerPostHandler) DeleteCommentPost(c *gin.Context) {
	userIdany, ok := c.Get("id")
	if !ok {
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	var comment models.DeleteCommentPost
	userId := userIdany.(int)
	comment.UserId = userId
	if err := c.ShouldBindJSON(&comment); err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	postOk, err := jph.GRPC_Client.DeleteCommentPost(comment.PostId, comment.UserId, comment.CommentId)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, postOk, nil)
	c.JSON(http.StatusOK, successRes)

}

// AddLikePost adds a like to a post.
// @Summary Add a like to a post
// @Description Add a like to a post by providing post ID
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param post_id query string true "Post ID"
// @Success 200 {object} response.Response "Like added successfully"
// @Failure 400 {object} response.Response "Bad request: incorrect format of post ID"
// @Failure 500 {object} response.Response "Internal server error: failed to add like to post"
// @Router /jobseeker/posts/like [post]
func (jph *JobseekerPostHandler) AddLikePost(c *gin.Context) {
	userIdany, ok := c.Get("id")
	if !ok {
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	userId := userIdany.(int)
	postIdstr := c.Query("post_id")
	postId, err := strconv.Atoi(postIdstr)
	if err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format of post id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	postOk, err := jph.GRPC_Client.AddLikePost(postId, userId)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, postOk, nil)
	c.JSON(http.StatusOK, successRes)

}

// RemoveLikePost removes a like from a post.
// @Summary Remove a like from a post
// @Description Remove a like from a post by providing post ID
// @Tags Jobseeker
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param Authorization header string true "Bearer token"
// @Param post_id query string true "Post ID"
// @Success 200 {object} response.Response "Like removed successfully"
// @Failure 400 {object} response.Response "Bad request: incorrect format of post ID"
// @Failure 500 {object} response.Response "Internal server error: failed to remove like from post"
// @Router /jobseeker/posts/like/remove [delete]
func (jph *JobseekerPostHandler) RemoveLikePost(c *gin.Context) {
	userIdany, ok := c.Get("id")
	if !ok {
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	userId := userIdany.(int)
	postIdstr := c.Query("post_id")
	postId, err := strconv.Atoi(postIdstr)
	if err != nil {
		errResp := response.ClientResponse(http.StatusBadRequest, "Incorrect Format of post id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	postOk, err := jph.GRPC_Client.RemoveLikePost(postId, userId)
	if err != nil {
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, msg.MsgSuccess, postOk, nil)
	c.JSON(http.StatusOK, successRes)

}
