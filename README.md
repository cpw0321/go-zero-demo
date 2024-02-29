# go-zero-demo

```shell
# execute 
goctl api go -api ./etc/api.api -dir . -style gozero

# generate model file
# psql
goctl model pg datasource --url="postgres://postgres:123456@localhost:5432/postgres?sslmode=disable" --table="*" -d ./model

# mysql
# sql
goctl model mysql ddl -c -src init.sql -dir ./model
# datasource
goctl model mysql datasource --url="root:123456@tcp(localhost:3306)/数据库名" -table 表名 -dir 目录路径

# add model logic
# seriesmodel.go

# run 
go run api.go
```