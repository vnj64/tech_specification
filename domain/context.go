package domain

type Context interface {
	MakeWithSession(s Session) Context

	Session() Session

	Services() Services
	Connection() Connection
}

func ValidateContext(c Context) error {
	return nil
}
