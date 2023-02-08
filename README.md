go install github.com/cloudwego/thrift-gen-validator@latest
# 生成服务代码or更新
kitex --thrift-plugin validator -module usermicro -service usermicro ./idl/user.thrift
kitex -module usermicro -service usermicro ./idl/user.thrift

go get github.com/cloudwego/kitex@latest && go mod tidy

# gorm相关代码
generateSchema ：生成sql schema
gormGen ：自动生成query代码

# 编写model
model/user.go

# 生成schema
cd cmd/generateSchema/ && go run gen.go

# 生成DAO
cd cmd/gormGen/ && go run gen.go