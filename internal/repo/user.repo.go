package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// ur: user repo
func (ur *UserRepo) GetInforUser() string {
	return "GetInforUser"
}
