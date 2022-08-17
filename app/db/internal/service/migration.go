package service

import (
    "context"

    pb "startkit/api/db/migration/v1"
)

type MigrationService struct {
    pb.UnimplementedMigrationServer
}

func NewMigrationService() *MigrationService {
    return &MigrationService{}
}

func (s *MigrationService) ListMigration(ctx context.Context, req *pb.ListMigrationRequest) (*pb.ListMigrationReply, error) {
    return &pb.ListMigrationReply{}, nil
}
