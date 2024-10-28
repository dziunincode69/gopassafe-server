package application

type UserService struct {
}

func (u *UserService) GetUser(username string) string {
	return "Hello " + username
}
