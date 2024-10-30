package port

type UserServicePort interface {
	Authenticate(email, password string) string
	Register(email, password string) string
}
