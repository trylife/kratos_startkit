package biz

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
)

type User struct {
    Id       int
    Uuid     string
    Nickname string
}

type UserRepo interface {
    GetUser(ctx context.Context, uuid string) (*User, error)
}

type UserUseCase struct {
    repo UserRepo
    log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
    return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/User"))}
}

func (uc *UserUseCase) Get(ctx context.Context, uuid string) (*User, error) {
    return uc.repo.GetUser(ctx, uuid)
}
