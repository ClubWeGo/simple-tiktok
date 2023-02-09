# user micro server

> 依赖 validator插件 ：`go install github.com/cloudwego/thrift-gen-validator@latest`
## 生成or更新服务代码
`kitex --thrift-plugin validator -module usermicro -service usermicro ./idl/user.thrift`

without validator : `kitex -module usermicro -service usermicro ./idl/user.thrift`

go get github.com/cloudwego/kitex@latest && go mod tidy

## gorm相关
- cmd/generateSchema ：生成sql schema
- cmd/gormGen ：自动生成query代码

## dal层
### model
./dal/model/user.go

### 生成schema
cd cmd/migrateSchema/ && go run migrate.go

### 生成query
cd cmd/gormGen && go run gen.go

### 创建service层
包装dal中的query，向handler暴露方法 

## 运行
- 编译 : sh ./build.sh
- 运行服务 : sh ./bootstrap.sh
- 运行client : cd /cmd/client && go run userclient.go

配置信息：

mysql8
- 库名 simpletk
- user tk
- password 123456
- dsn := "tk:123456@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"

etcd
- 0.0.0.0:2379

docker部署
- `docker build -t imagename:version .`
- `docker run -t -i --rm -p 8888:8888 --name testserver imagename:version /bin/bash`