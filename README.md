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

# TODO lists
- [x] 基础功能
- [x] 整合monio的publish视频功能 : 写法和性能还需要优化

    minio:  `docker run -it -p 9000:9000 -p 43543:43543 -d minio/minio server /data --console-address ":43543" --address ":9000"`
    api: `127.0.0.1:9000`
    console: `127.0.0.1:43543`
    
    创建douyin桶，
    设置douyin桶策略的权限为public

- [ ] feed目前是分页的方式，需要改为题目要求的时间点+limit的方式，videomicro需要同步改动
- [ ] 需要加vd校验字段，重新生成带校验的代码
- [ ] token登录(jwt?)，注意视频上传需要通过token查user，实现token的时候记得实现tools token中的ValidateToken返回对应userid
- [ ] 可选功能，选哪个？，(基础部分的isfollow与isFavourite)
- [ ] ...


依赖另外两个微服务的代码仓库

go get github.com/ClubWeGo/videomicro@latest

go get github.com/ClubWeGo/usermicro@latest