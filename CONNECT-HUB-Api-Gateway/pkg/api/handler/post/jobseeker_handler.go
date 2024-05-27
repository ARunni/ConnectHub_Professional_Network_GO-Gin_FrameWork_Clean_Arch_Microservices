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
