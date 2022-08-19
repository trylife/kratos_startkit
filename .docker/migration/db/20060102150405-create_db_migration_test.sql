-- init database when docker-compose environment setup

-- +migrate Up
create database if not exists migration_test;

-- +migrate Down
drop database if exists migration_test;