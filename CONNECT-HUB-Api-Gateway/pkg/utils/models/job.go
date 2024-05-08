
package models

import (
	"time"
)

type JobOpening struct {
	ID                  uint      `json:"id"`
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

type JobOpeningResponse struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	Requirements        string    `json:"requirements"`
	PostedOn            time.Time `json:"posted_on"`
	RecruiterID          int32     `json:"recruiter_id"`
	Location            string    `json:"location"`
	EmploymentType      string    `json:"employment_type"`
	Salary              string    `json:"salary"`
	SkillsRequired      string    `json:"skills_required"`
	ExperienceLevel     string    `json:"experience_level"`
	EducationLevel      string    `json:"education_level"`
	ApplicationDeadline time.Time `json:"application_deadline"`
}

type AllJob struct {
	ID                  uint      `json:"id"`
	Title               string    `json:"title"`
	ApplicationDeadline time.Time `json:"application_deadline"`
	RecruiterID          int32     `json:"recruiter_id"`
}
