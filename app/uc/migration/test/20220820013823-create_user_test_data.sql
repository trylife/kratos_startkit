-- +migrate Up
insert into user(id, nickname, uuid, phone_number, create_time, update_time)
values (1, 'John Smith', 'TEST_FIRST_UUID', '18566668888', unix_timestamp(), unix_timestamp());

-- +migrate Down
delete
from user
where uuid = 'TEST_FIRST_UUID';