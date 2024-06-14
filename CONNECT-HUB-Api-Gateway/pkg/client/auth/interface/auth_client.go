package interfaces

type AuthClient interface {
	VideoCallKey(userID, oppositeUser int) (string, error)
}
