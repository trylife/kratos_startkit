-- init database when docker-compose environment setup

-- +migrate Up
create database migration_test;

-- +migrate Down
drop database migration_test;