### 环境配置及依赖安装
```shell
export GOPATH=~/go
export PATH=$GOPATH/bin:$PATH
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
go install github.com/cloudwego/thriftgo@latest
```

### 生成或更新 kitex 脚手架代码
```shell
kitex -module github.com/ClubWeGo/favoritemicro -service favoritemicro idl/favoritemicro.thrift
```
