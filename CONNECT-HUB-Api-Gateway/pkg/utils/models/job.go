package models

import "time"

type JobOpening struct {
	EmployerID          int       `json:"employer_id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}

type JobOpeningReq struct {

	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}

type JobOpeningData struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	PostedOn            time.Time `json:"posted_on"`
	EmployerID          int       `json:"employer_id"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              int       `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}

type AllJob struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	EmployerID          int32     `json:"employer_id"`
}

type JobSeekerGetAllJobs struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type AppliedJobs struct {
	Jobs []ApplyJobRes `json:"jobs"`
}

type AppliedJobsJ struct {
	Jobs []ApplyJob `json:"jobs"`
}

type ApplyJobReq struct {
	JobID       uint   `json:"job_id"`
	JobseekerID uint   `json:"jobseeker_id"`
	CoverLetter string `json:"cover_letter"`
	Resume      []byte `json:"resume"`
}

type ApplyJob struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	JobID       uint   `json:"job_id"`
	JobseekerID uint   `json:"jobseeker_id"`
	RecruiterID uint   `json:"recruiter_id"`
	CoverLetter string `json:"cover_letter"`
	ResumeUrl   string `json:"resume_url"`
	Status      string `gorm:"default:waiting" json:"status"`
}
type ApplyJobRes struct {
	ID            uint   `gorm:"primary_key;auto_increment" json:"id"`
	JobID         uint   `json:"job_id"`
	JobseekerID   uint   `json:"jobseeker_id"`
	JobseekerName string `json:"jobseeker_name"`
	JoseekerEmail string `json:"jobseeker_email"`
	RecruiterID   uint   `json:"recruiter_id"`
	CoverLetter   string `json:"cover_letter"`
	ResumeUrl     string `json:"resume_url"`
	Status        string `gorm:"default:waiting" json:"status"`
}

type Interview struct {
	ID            uint      `gorm:"primary_key;auto_increment" json:"id"`
	ApplicationId uint      `json:"application_id"`
	JobID         uint      `json:"job_id"`
	JobseekerID   uint      `json:"jobseeker_id"`
	RecruiterID   uint      `json:"recruiter_id"`
	DateAndTime   time.Time `json:"date_and_time"`
	Mode          string    `gorm:"default:online" json:"mode"`
	Link          string    `json:"link"`
	Status        string    `gorm:"default:scheduled" json:"status"`
}

type ScheduleReq struct {
	ApplicationId int       `json:"application_id"`
	RecruiterID   uint      `json:"recruiter_id"`
	DateAndTime   time.Time `json:"date_and_time"`
	Mode          string    `gorm:"default:online" json:"mode"`
	Link          string    `json:"link"`
}

type ScheduleReqs struct {
	ApplicationId int       `json:"application_id"`
	DateAndTime   time.Time `json:"date_and_time"`
	Mode          string    `gorm:"default:online" json:"mode"`
	Link          string    `json:"link"`
}
