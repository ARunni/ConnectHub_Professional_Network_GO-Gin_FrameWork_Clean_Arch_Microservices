package repository

import (
	"ConnetHub_post/pkg/repository/interfaces"
	"ConnetHub_post/pkg/utils/models"

	"gorm.io/gorm"
)

type jobseekerPostRepository struct {
	DB *gorm.DB
}

func NewJobseekerPostRepository(DB *gorm.DB) interfaces.JobseekerPostRepository {
	return &jobseekerPostRepository{
		DB: DB,
	}
}

func (jr *jobseekerPostRepository) CreatePost(post models.CreatePostRes) (models.CreatePostRes, error) {
	var res models.CreatePostRes
	querry := `
		insert into table posts 
		(title,content,image_url,jobseeker_id) 
		values (?,?,?,?,?) returnes *
		`

	err := jr.DB.Exec(querry, post.Title, post.Content, post.ImageUrl, post.JobseekerId).Scan(&res).Error
	if err != nil {
		return models.CreatePostRes{}, err
	}
	return res, nil
}

func (jr *jobseekerPostRepository) GetOnePost(postId int) (models.CreatePostRes, error) {
	var res models.CreatePostRes
	querry := `
		select id,jobseeker_id,title,content,image_url,created_at 
		from posts where id = ?
		`

	err := jr.DB.Raw(querry, postId).Scan(&res).Error
	if err != nil {
		return models.CreatePostRes{}, err
	}
	return res, nil
}

func (jr *jobseekerPostRepository) GetAllPost() (models.AllPost, error) {
	var res models.AllPost
	querry := `
		select id,jobseeker_id,title,content,image_url,created_at 
		from posts
		`

	err := jr.DB.Raw(querry).Scan(&res).Error
	if err != nil {
		return models.AllPost{}, err
	}
	return res, nil
}

func (jr *jobseekerPostRepository) UpdatePost(post models.EditPostRes) (models.EditPostRes, error) {
	var res models.EditPostRes
	querry := `
		update table posts title = ?,content = ?,image_url ?,created_at =? 
		from posts where id = ? and jobseeker_id = ?
		`

	err := jr.DB.Exec(querry, post.Title, post.Content, post.ImageUrl).Scan(&res).Error
	if err != nil {
		return models.EditPostRes{}, err
	}
	return res, nil
}

func (jr *jobseekerPostRepository) DeletePost(postId, JobseekerId int) (bool, error) {

	querry := `
	delete from posts 
	where id = ? and jobseeker_id = ?
	`
	err := jr.DB.Exec(querry, postId, JobseekerId).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (jr *jobseekerPostRepository) IsPostExistByPostId(postId int) (bool, error) {
	var ok int
	querry := `
	select count(*) from posts 
	where id = ?
	`
	err := jr.DB.Raw(querry, postId).Scan(&ok).Error
	if err != nil {
		return false, err
	}
	return ok > 0, nil
}

func (jr *jobseekerPostRepository) IsPostExistByUserId(userId int) (bool, error) {
	var ok int
	querry := `
	select count(*) from posts 
	where jobseeker_id = ?
	`
	err := jr.DB.Raw(querry, userId).Scan(&ok).Error
	if err != nil {
		return false, err
	}
	return ok > 0, nil
}
