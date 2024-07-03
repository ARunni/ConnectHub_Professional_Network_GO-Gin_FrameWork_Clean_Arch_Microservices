package handler

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	logging "github.com/ARunni/connectHub_gateway/Logging"
	"github.com/ARunni/connectHub_gateway/pkg/client/post/interfaces"
	"github.com/ARunni/connectHub_gateway/pkg/utils/models"
	"github.com/ARunni/connectHub_gateway/pkg/utils/response"

	msg "github.com/ARunni/Error_Message"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type JobseekerPostHandler struct {
	GRPC_Client interfaces.JobseekerPostClient
	Logger      *logrus.Logger
	LogFile     *os.File
}

func NewJobseekerPostHandler(grpc_client interfaces.JobseekerPostClient) *JobseekerPostHandler {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_gateway.log")
	return &JobseekerPostHandler{
		GRPC_Client: grpc_client,
		Logger:      logger,
		LogFile:     logFile,
	}
}

// CreatePost creates a new post by a jobseeker.
// @Summary Create a post
// @Description Create a new post by a jobseeker
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
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
		jph.Logger.Error("Error on Getting Data", errors.New("id error"))
		err := errors.New(msg.ErrGetData)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	jobseekerId := jobseekerIdAny.(int)
	post.JobseekerId = jobseekerId

	file, err := c.FormFile("image")
	if err != nil {
		jph.Logger.Error("Error on Getting image", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Getting Image Failed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	fileContent, err := file.Open()
	if err != nil {
		jph.Logger.Error("Error on opennig image", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Error opening the file", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return

	}
	defer fileContent.Close()
	imageData, err := ioutil.ReadAll(fileContent)
	if err != nil {
		jph.Logger.Error("Error on reading image", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, "Error reading the file", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	post.Image = imageData

	postData, err := jph.GRPC_Client.CreatePost(post)

	if err != nil {
		jph.Logger.Error("Error on Create Post", err)
		errResp := response.ClientResponse(http.StatusInternalServerError, msg.ErrInternal, nil, err.Error())
		c.JSON(http.StatusInternalServerError, errResp)
		return
	}
	jph.Logger.Info("Create Post successful")
	successRes := response.ClientResponse(http.StatusOK, msg.MsgAddSuccess, postData, nil)
	c.JSON(http.StatusOK, successRes)

}

// GetOnePost retrieves a single post by its ID.
// @Summary Get a Post by ID
// @Description Retrieve a single post by its ID
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
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


// GetAllPost retrieves all posts for jobseekers.
// @Summary Get all posts
// @Description Retrieves all posts for jobseekers
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully retrieved all posts"
// @Failure 500 {object} response.Response "Internal server error: failed to retrieve posts"
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


// UpdatePost updates an existing post for a jobseeker.
// @Summary Update post
// @Description Updates an existing post for a jobseeker
// @Tags Jobseeker Post Management
// @Accept multipart/form-data
// @Produce json
// @Security BearerTokenAuth
// @Param post_id formData int true "ID of the post to update"
// @Param title formData string true "Title of the post"
// @Param content formData string true "Content of the post"
// @Param image formData file true "Image for the post"
// @Success 200 {object} response.Response "Successfully updated the post"
// @Failure 400 {object} response.Response "Bad request: invalid input parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to update post"
// @Router /jobseeker/post [patch]
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

// DeletePost deletes an existing post for a jobseeker.
// @Summary Delete post
// @Description Deletes an existing post for a jobseeker
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param post_id query int true "ID of the post to delete"
// @Success 200 {object} response.Response "Successfully deleted the post"
// @Failure 400 {object} response.Response "Bad request: invalid input parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to delete post"
// @Router /jobseeker/post [delete]
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


// CreateCommentPost adds a comment to an existing post for a jobseeker.
// @Summary Create comment on post
// @Description Adds a comment to an existing post for a jobseeker
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param comment body models.CreateCommentPostReq true "Comment details"
// @Success 200 {object} response.Response "Successfully added comment to the post"
// @Failure 400 {object} response.Response "Bad request: invalid input parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to add comment to the post"
// @Router /jobseeker/post/comment [post]
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

// UpdateCommentPost updates a comment on an existing post for a jobseeker.
// @Summary Update comment on post
// @Description Updates a comment on an existing post for a jobseeker
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param comment body models.UpdateCommentPostReq true "Updated comment details"
// @Success 200 {object} response.Response "Successfully updated comment on the post"
// @Failure 400 {object} response.Response "Bad request: invalid input parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to update comment on the post"
// @Router /jobseeker/post/comment [put]
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


// DeleteCommentPost deletes a comment on an existing post for a jobseeker.
// @Summary Delete comment on post
// @Description Deletes a comment on an existing post for a jobseeker
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param comment body models.DeleteCommentPostReq true "Comment deletion details"
// @Success 200 {object} response.Response "Successfully deleted comment from the post"
// @Failure 400 {object} response.Response "Bad request: invalid input parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to delete comment from the post"
// @Router /jobseeker/post/comment [delete]
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


// AddLikePost adds a like to a post for a jobseeker.
// @Summary Add like to post
// @Description Adds a like to a post for a jobseeker
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param post_id query int true "ID of the post to like"
// @Success 200 {object} response.Response "Successfully added like to the post"
// @Failure 400 {object} response.Response "Bad request: invalid input parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to add like to the post"
// @Router /jobseeker/post/like [post]
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


// RemoveLikePost removes a like from a post for a jobseeker.
// @Summary Remove like from post
// @Description Removes a like from a post for a jobseeker
// @Tags Jobseeker Post Management
// @Accept json
// @Produce json
// @Security BearerTokenAuth
// @Param post_id query int true "ID of the post to remove like"
// @Success 200 {object} response.Response "Successfully removed like from the post"
// @Failure 400 {object} response.Response "Bad request: invalid input parameters"
// @Failure 500 {object} response.Response "Internal server error: failed to remove like from the post"
// @Router /jobseeker/post/like [delete]
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
