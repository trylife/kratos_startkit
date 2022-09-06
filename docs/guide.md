# CMD


## components

### migrate

```bash
brew install protobuf
brew install protoc-gen-go
go get -v github.com/rubenv/sql-migrate/...
go install -v github.com/rubenv/sql-migrate/...

docker-compose up -d 

# 创建数据库
sql-migrate up -config=configs/migration.yaml -env="development_init"

# 初始化数据结构与基础数据
sql-migrate up -config=configs/migration.yaml -env="development_basic"

# 载入测试数据
sql-migrate up -config=configs/migration.yaml -env="development_test"
```

# References
- http://localhost:10000/helloworld/kratos
- DB Ent https://entgo.io/docs/getting-started/
- Migrate https://github.com/rubenv/sql-migrate


```bash
kratos proto add api/uc/user/v1/user.proto 
kratos proto client api/uc/user/v1/user.proto   
kratos proto server api/uc/user/v1/user.proto -t internal/service

```


# structs


```
/uc/v1/user/:user_uuid
/{appname}/{module}/{version}/actions 
```