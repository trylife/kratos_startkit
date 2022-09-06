-- +migrate Up
create table if not exists user
(
    id           int(11) unsigned   not null auto_increment primary key comment 'id',
    nickname     varchar(32)        not null,
    uuid         varchar(64)        not null,
    phone_number varchar(16) UNIQUE not null,
    create_time  int(11)            not null default 0,
    update_time  int(11)            not null default 0
);

-- +migrate Down
drop table if exists user;