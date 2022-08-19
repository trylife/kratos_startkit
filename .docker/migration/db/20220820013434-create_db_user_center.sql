
-- +migrate Up
create database if not exists user_center;
-- +migrate Down
drop database if exists user_center;
