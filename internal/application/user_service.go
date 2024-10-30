package application

type UserService struct {
}

func (u *UserService) Authenticate(email, password string) string {
	return "0000000000-LOGIN-00000-00000"
}
func (u *UserService) Register(email, password string) string {
	return "0000000000-REGISTER-00000-00000"
}
