package data

import (
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
    "startkit/app/uc/internal/conf"
    "startkit/app/uc/internal/data/ent"

    // init mysql driver
    _ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
    NewData,
    NewEntClient,
    NewUserRepo,
)

// Data .
type Data struct {
    db  *ent.Client
    log *log.Helper
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
    log := log.NewHelper(log.With(logger, "module", "uc/data/ent"))

    client, err := ent.Open(
        conf.Database.Driver,
        conf.Database.Source,
    )
    if err != nil {
        log.Fatalf("failed opening connection to db: %v", err)
    }
    // Run the auto migration tool.
    //if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
    //    log.Fatalf("failed creating schema resources: %v", err)
    //}
    return client
}

// NewData .
// NewData .
func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
    log := log.NewHelper(log.With(logger, "module", "uc/data"))

    d := &Data{
        db:  entClient,
        log: log,
    }
    return d, func() {
        if err := d.db.Close(); err != nil {
            log.Error(err)
        }
    }, nil
}
