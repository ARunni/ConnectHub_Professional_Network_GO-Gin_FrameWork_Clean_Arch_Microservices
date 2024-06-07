package interfaces

type JobAuthClient interface {
	GetDetailsById(userId int) (string, string, error)
}