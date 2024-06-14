package interfaces

type VideoCallUsecase interface {
	VideoCallKey(userID, oppositeUser int) (string, error)
}