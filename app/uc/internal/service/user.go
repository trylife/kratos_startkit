package service

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    pb "startkit/api/uc/user/v1"
    "startkit/app/uc/internal/biz"
)

type UserService struct {
    pb.UnimplementedUserServer

    uc  *biz.UserUseCase
    log *log.Helper
}

func NewUserService(useCase *biz.UserUseCase, logger log.Logger) *UserService {
    return &UserService{
        uc:  useCase,
        log: log.NewHelper(log.With(logger, "module", "user/catalog")),
    }
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
    return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
    return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
    return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
    u, err := s.uc.Get(ctx, req.Uuid)
    if err != nil {
        return nil, err
    }
    return &pb.GetUserReply{Uuid: u.Uuid, Nickname: u.Nickname}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
    return &pb.ListUserReply{}, nil
}
