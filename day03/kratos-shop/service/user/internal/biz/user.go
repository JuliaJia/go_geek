package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID       int64
	Mobile   string
	Password string
	NickName string
	Birthday int64
	Gender   string
	Role     int
}

//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
type UserRepo interface {
	CreateUser(context.Context, *User) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, u *User) (*User, error) {
	return uc.repo.CreateUser(ctx, u)
}
