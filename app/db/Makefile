include ../../app_makefile


# additional commands
.PHONY: migrate
migrate:
	# docker-compose up -d
	sql-migrate up -config=configs/migration.yaml -env="development_init"
	sql-migrate up -config=configs/migration.yaml -env="development_basic"
	sql-migrate up -config=configs/migration.yaml -env="development_test"