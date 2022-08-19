package server

import (
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    v1 "startkit/api/db/migration/v1"
    "startkit/app/db/internal/conf"
    "startkit/app/db/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
    c *conf.Server,
// greeter *service.GreeterService,
    migration *service.MigrationService,
    logger log.Logger,
) *grpc.Server {
    var opts = []grpc.ServerOption{
        grpc.Middleware(
            recovery.Recovery(),
        ),
    }
    if c.Grpc.Network != "" {
        opts = append(opts, grpc.Network(c.Grpc.Network))
    }
    if c.Grpc.Addr != "" {
        opts = append(opts, grpc.Address(c.Grpc.Addr))
    }
    if c.Grpc.Timeout != nil {
        opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
    }
    srv := grpc.NewServer(opts...)
    // v1.RegisterGreeterServer(srv, greeter)
    v1.RegisterMigrationServer(srv, migration)
    return srv
}
