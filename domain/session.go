package domain

type Session interface {
	Authorized() bool
	UserId() int
}
