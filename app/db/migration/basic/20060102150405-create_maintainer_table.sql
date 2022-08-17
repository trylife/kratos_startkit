-- +migrate Up
CREATE TABLE IF NOT EXISTS project_maintainer
(
    id                  int(11) unsigned not null auto_increment primary key comment 'id',
    resource_type       varchar(16)      not null comment 'resource_type i.e. database, table, rest_api, grpc_api',
    resource_name       varchar(32)      NOT NULL comment 'i.e. erp.resource_owner',
    resource_owner      json             not null comment 'resource_owner array',
    resource_maintainer json             not null comment 'resource_maintainer ',
    create_time         int(11)          NOT NULL default 0,
    update_time         int(11)          NOT NULL default 0
);

-- +migrate Down
DROP TABLE IF EXISTS resourse_owner;