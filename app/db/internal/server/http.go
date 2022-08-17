package server

import (
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/middleware/recovery"
    "github.com/go-kratos/kratos/v2/transport/http"
    v1 "startkit/api/db/migration/v1"
    "startkit/app/db/internal/conf"
    "startkit/app/db/internal/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
    c *conf.Server,
//greeter *service.GreeterService,
    migration *service.MigrationService,
    logger log.Logger,
) *http.Server {
    var opts = []http.ServerOption{
        http.Middleware(
            recovery.Recovery(),
        ),
    }
    if c.Http.Network != "" {
        opts = append(opts, http.Network(c.Http.Network))
    }
    if c.Http.Addr != "" {
        opts = append(opts, http.Address(c.Http.Addr))
    }
    if c.Http.Timeout != nil {
        opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
    }
    srv := http.NewServer(opts...)
    // v1.RegisterGreeterHTTPServer(srv, greeter)
    v1.RegisterMigrationHTTPServer(srv, migration)
    return srv
}