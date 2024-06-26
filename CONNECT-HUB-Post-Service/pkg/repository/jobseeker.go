package repository

import (
	"os"
	"time"

	logging "github.com/ARunni/ConnetHub_post/Logging"
	"github.com/ARunni/ConnetHub_post/pkg/repository/interfaces"
	"github.com/ARunni/ConnetHub_post/pkg/utils/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type jobseekerPostRepository struct {
	DB      *gorm.DB
	Logger  *logrus.Logger
	LogFile *os.File
}

func NewJobseekerPostRepository(DB *gorm.DB) interfaces.JobseekerPostRepository {
	logger, logFile := logging.InitLogrusLogger("./Logging/connectHub_Post.log")
	return &jobseekerPostRepository{
		DB:      DB,
		Logger:  logger,
		LogFile: logFile,
	}
}

func (jr *jobseekerPostRepository) CreatePost(post models.CreatePostRes) (models.CreatePostRes, error) {
	jr.Logger.Info("CreatePost at jobseekerPostRepository started")

	var res models.CreatePostRes
	query := `
		INSERT INTO posts (title, content, image_url, jobseeker_id,created_at) 
		VALUES ($1, $2, $3, $4,$5) 
		RETURNING id, title, content, image_url, jobseeker_id, created_at
	`

	// Using Raw SQL and Scan to execute and get the returning values
	row := jr.DB.Raw(query, post.Title, post.Content, post.ImageUrl, post.JobseekerId, time.Now()).Row()
	err := row.Scan(&res.ID, &res.Title, &res.Content, &res.ImageUrl, &res.JobseekerId, &res.CreatedAt)
	if err != nil {
		jr.Logger.Error("error at database", err)
		return models.CreatePostRes{}, err
	}
	jr.Logger.Info("CreatePost at jobseekerPostRepository finished")
	return res, nil
}

func (jr *jobseekerPostRepository) GetOnePost(postId int) (models.CreatePostResp, error) {
	jr.Logger.Info("GetOnePost at jobseekerPostRepository started")
	var res models.CreatePostResp
	querry := `
		select id,jobseeker_id,title,content,image_url,created_at 
		from posts where id = ?
		`

	err := jr.DB.Raw(querry, postId).Scan(&res).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return models.CreatePostResp{}, err
	}
	jr.Logger.Info("GetOnePost at jobseekerPostRepository finished")
	return res, nil
}

func (jr *jobseekerPostRepository) GetAllPost() (models.AllPost, error) {
	jr.Logger.Info("GetAllPost at jobseekerPostRepository started")
	var res []models.CreatePostResp
	querry := `
		select id,jobseeker_id,title,content,image_url,created_at
		from posts
		`

	err := jr.DB.Raw(querry).Scan(&res).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return models.AllPost{}, err
	}
	allPosts := models.AllPost{
		Posts: res,
	}
	jr.Logger.Info("GetAllPost at jobseekerPostRepository finished")
	return allPosts, nil
}

//

func (jr *jobseekerPostRepository) UpdatePost(post models.EditPostRes) (models.EditPostRes, error) {
	jr.Logger.Info("UpdatePost at jobseekerPostRepository started")
	var res models.EditPostRes
	querry := `
		update posts set title = $1,content = $2,image_url = $3,created_at = $4 
		 where id = $5 and jobseeker_id = $6 returning id as post_id, jobseeker_id, title,content,image_url,updated_at
		`

	err := jr.DB.Raw(querry, post.Title, post.Content, post.ImageUrl, time.Now(), post.PostId, post.JobseekerId).Scan(&res).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return models.EditPostRes{}, err
	}
	jr.Logger.Info("UpdatePost at jobseekerPostRepository finished")
	return res, nil
}

func (jr *jobseekerPostRepository) DeletePost(postId, JobseekerId int) (bool, error) {
	jr.Logger.Info("DeletePost at jobseekerPostRepository started")
	querry := `
	delete from posts 
	where id = ? and jobseeker_id = ?
	`
	err := jr.DB.Exec(querry, postId, JobseekerId).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("DeletePost at jobseekerPostRepository finished")
	return true, nil
}

func (jr *jobseekerPostRepository) IsPostExistByPostId(postId int) (bool, error) {
	jr.Logger.Info("IsPostExistByPostId at jobseekerPostRepository started")
	var ok int
	querry := `
	select count(*) from posts 
	where id = ?
	`
	err := jr.DB.Raw(querry, postId).Scan(&ok).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("IsPostExistByPostId at jobseekerPostRepository finished")
	return ok > 0, nil
}

func (jr *jobseekerPostRepository) IsPostExistByUserId(userId int) (bool, error) {
	jr.Logger.Info("IsPostExistByUserId at jobseekerPostRepository started")
	var ok int
	querry := `
	select count(*) from posts 
	where jobseeker_id = ?
	`
	err := jr.DB.Raw(querry, userId).Scan(&ok).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("IsPostExistByUserId at jobseekerPostRepository finished")
	return ok > 0, nil
}

func (jr *jobseekerPostRepository) CreateCommentPost(postId, userId int, comment string) (bool, error) {
	jr.Logger.Info("CreateCommentPost at jobseekerPostRepository started")
	querry := `insert into comments (post_id,comment,jobseeker_id,created_at)
	values ($1,$2,$3,$4)`
	err := jr.DB.Exec(querry, postId, comment, userId, time.Now()).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("CreateCommentPost at jobseekerPostRepository finished")
	return true, nil
}

func (jr *jobseekerPostRepository) IsCommentIdExist(commentId int) (bool, error) {
	jr.Logger.Info("IsCommentIdExist at jobseekerPostRepository started")
	var ok int
	querry := `select count(*) from comments where id = ?`
	err := jr.DB.Raw(querry, commentId).Scan(&ok).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("IsCommentIdExist at jobseekerPostRepository finished")
	return ok > 0, nil
}

func (jr *jobseekerPostRepository) IsCommentIdBelongsUserId(commentId, userId int) (bool, error) {
	jr.Logger.Info("IsCommentIdBelongsUserId at jobseekerPostRepository started")
	var ok int
	querry := `select count(*) from comments where id = ? and  jobseeker_id = ?`
	err := jr.DB.Raw(querry, commentId, userId).Scan(&ok).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("IsCommentIdBelongsUserId at jobseekerPostRepository finished")
	return ok > 0, nil
}

func (jr *jobseekerPostRepository) UpdateCommentPost(commentId, postId, userId int, comment string) (bool, error) {
	jr.Logger.Info("UpdateCommentPost at jobseekerPostRepository started")
	querry := `update comments set comment = ?, updated_at = ? where id = ? and  jobseeker_id = ? and post_id = ? `
	err := jr.DB.Exec(querry, comment, time.Now(), commentId, userId, postId).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("UpdateCommentPost at jobseekerPostRepository finished")
	return true, nil
}

func (jr *jobseekerPostRepository) DeleteCommentPost(postId, userId, commentId int) (bool, error) {
	jr.Logger.Info("DeleteCommentPost at jobseekerPostRepository started")
	querry := `delete from comments where id = ? and  jobseeker_id = ? and post_id = ? `
	err := jr.DB.Exec(querry, commentId, userId, postId).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("DeleteCommentPost at jobseekerPostRepository finished")
	return true, nil
}

func (jr *jobseekerPostRepository) IsLikeExist(postId, userId int) (bool, error) {
	jr.Logger.Info("IsLikeExist at jobseekerPostRepository started")
	var ok int
	querry := `select count(*) from likes where jobseeker_id = ? and post_id = ? `
	err := jr.DB.Raw(querry, userId, postId).Scan(&ok).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("IsLikeExist at jobseekerPostRepository finished")
	return ok > 0, nil
}

func (jr *jobseekerPostRepository) AddLikePost(postId, userId int) (bool, error) {
	jr.Logger.Info("AddLikePost at jobseekerPostRepository started")
	querry := `insert into likes (jobseeker_id,post_id,created_at) values (?,?,?)`
	err := jr.DB.Exec(querry, userId, postId, time.Now()).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("AddLikePost at jobseekerPostRepository finished")
	return true, nil
}

func (jr *jobseekerPostRepository) RemoveLikePost(postId, userId int) (bool, error) {
	jr.Logger.Info("RemoveLikePost at jobseekerPostRepository started")
	querry := `delete from likes where post_id = ? and  jobseeker_id = ?`
	err := jr.DB.Exec(querry, postId, userId).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return false, err
	}
	jr.Logger.Info("RemoveLikePost at jobseekerPostRepository finished")
	return true, nil
}

func (jr *jobseekerPostRepository) GetCommentsPost(postId int) ([]models.CommentData, error) {
	jr.Logger.Info("GetCommentsPost at jobseekerPostRepository started")

	var comments []models.CommentData

	querry := `select id,comment,jobseeker_id,created_at,updated_at
	from comments where post_id = ?`
	err := jr.DB.Raw(querry, postId).Scan(&comments).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return nil, err
	}
	jr.Logger.Info("GetCommentsPost at jobseekerPostRepository finished")
	return comments, nil
}

func (jr *jobseekerPostRepository) GetLikesCountPost(postId int) (int, error) {
	jr.Logger.Info("GetLikesCountPost at jobseekerPostRepository started")
	var count int
	query := `SELECT count(*) 
	FROM likes WHERE post_id = ?`
	err := jr.DB.Raw(query, postId).Scan(&count).Error
	if err != nil {
		jr.Logger.Error("error at database", err)
		return 0, err
	}
	jr.Logger.Info("GetLikesCountPost at jobseekerPostRepository finished")
	return count, nil
}
