package interfaces

type VideoCallRepository interface {
	IsJobseekerExist(userId int) (bool, error)
	IsRecruiterExist(userId int) (bool, error)
}
