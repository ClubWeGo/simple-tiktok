<<<<<<< HEAD
## 生成代码
`hz new -idl idl/core.thrift -module github.com/ClubWeGo/douyin`

> 扩展模块

> `hz new -idl idl/interaction.thrift -module github.com/ClubWeGo/douyin`
> `hz new -idl idl/relation.thrift -module github.com/ClubWeGo/douyin`

## 删除代码
`rm -r ./biz router.go router_gen.go main.go go.mod go.sum .hz .gitignore`

## 更新代码
`hz update -idl idl/core.thrift`

## 拉取依赖
`go mod tidy`

## 编译代码
`go build`


依赖另外几个微服务的代码仓库

(微服务代码更新后，需主动执行来更新代码)

go get github.com/ClubWeGo/videomicro@latest

go get github.com/ClubWeGo/usermicro@latest

go get github.com/ClubWeGo/relationmicro@latest

go get github.com/ClubWeGo/favoritemicro@latest

go get github.com/ClubWeGo/commentmicro@latest


# 说明
1. 当前注册用户时，前端并未提供个人参数配置的能力，也没有提供更新用户信息的能力，所以在core.register_server处写死了初始化的用户背景图像和头像。在配置本项目时，请将对应文件名换成minio中存储的文件名。
=======
# user micro server

> 依赖 validator插件 ：`go install github.com/cloudwego/thrift-gen-validator@latest`
## 生成or更新服务代码
> 目前项目微服务不再做字段检查校验，这部分交给业务层单独实现安全机制去处理

`kitex --thrift-plugin validator -module github.com/ClubWeGo/usermicro -service usermicro ./idl/user.thrift`

without validator : `kitex -module github.com/ClubWeGo/usermicro -service usermicro ./idl/user.thrift`

注意，一定要先go get下kitex：`go get github.com/cloudwego/kitex@latest && go mod tidy`

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
>>>>>>> dev
