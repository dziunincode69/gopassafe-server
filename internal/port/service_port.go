package port

type UserServicePort interface {
	GetUser(username string) string
}
