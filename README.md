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
