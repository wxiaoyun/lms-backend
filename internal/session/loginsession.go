package session

type LoginSession struct {
	UserID         uint
	Email          string
	IsMasquerading bool
}
