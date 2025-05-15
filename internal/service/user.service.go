package service

import "github.com/HoangTrungAnh2k4/Ecommerce_GO/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// us: user service
func (us *UserService) RegisterUser() string {
	return us.userRepo.GetInforUser()
}
