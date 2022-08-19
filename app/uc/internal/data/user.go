package data

import (
    "context"
    "github.com/go-kratos/kratos/v2/log"
    "startkit/app/uc/internal/biz"
    "startkit/app/uc/internal/data/ent/user"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
    data *Data
    log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
    return &userRepo{
        data: data,
        log:  log.NewHelper(log.With(logger, "module", "data/beer")),
    }
}

func (r *userRepo) GetUser(ctx context.Context, uuid string) (*biz.User, error) {
    u, err := r.data.db.User.
        Query().
        Where(user.UUID(uuid)).
        Only(ctx)
    if err != nil {
        return nil, err
    }
    return &biz.User{
        Id:       u.ID,
        Uuid:     u.UUID,
        Nickname: u.Nickname,
    }, nil
}
