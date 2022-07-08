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
	ListUser(context context.Context, pageNum, pageSize int) ([]*User, int, error)
	GetUserByMobile(context context.Context, mobile string) (*User, error)
	UpdateUser(context.Context, *User) (bool, error)
	CheckPassword(context context.Context, p, ep string) (bool, error)
	GetUserById(context context.Context, id int64) (*User, error)
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

func (uc *UserUsecase) ListUser(ctx context.Context, pageNum, pageSize int) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUsecase) GetUserByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.GetUserByMobile(ctx, mobile)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, u *User) (bool, error) {

	return uc.repo.UpdateUser(ctx, u)
}

func (uc *UserUsecase) CheckPassword(ctx context.Context, p, ep string) (bool, error) {

	return uc.repo.CheckPassword(ctx, p, ep)
}
func (uc *UserUsecase) GetUserById(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUserById(ctx, id)
}
