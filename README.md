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

# list
- [ ] 基础功能publish视频还没有实现
- [ ] vd校验字段
- [ ] token登录(是否用jwt?)
- [ ] 可选功能(基础部分的isfollow与isFavourite)
- [ ] ...


依赖另外两个微服务的代码仓库

go get github.com/ClubWeGo/videomicro@latest

go get github.com/ClubWeGo/usermicro@latest