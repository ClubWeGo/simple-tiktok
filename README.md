# user微服务

> 依赖 validator插件 ：go install github.com/cloudwego/thrift-gen-validator@latest
# 生成服务代码or更新
kitex --thrift-plugin validator -module usermicro -service usermicro ./idl/user.thrift
> without validator : kitex -module usermicro -service usermicro ./idl/user.thrift

go get github.com/cloudwego/kitex@latest && go mod tidy

# gorm相关代码
cmd/generateSchema ：生成sql schema
cmd/gormGen ：自动生成query代码

# 创建dal层
## 编写model
./dal/model/user.go
写入model

# 生成schema
cd cmd/migrateSchema/ && go run migrate.go

# 生成DAO
cd cmd/gormGen && go run gen.go

# 创建service层
包装dal中的query，向handler暴露方法 

# 测试
编译 : sh ./build.sh

运行服务 : sh ./bootstrap.sh

运行client : cd /cmd/client && go run userclient.go