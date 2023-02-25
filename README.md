# simple-tiktok

基于 kitex RPC微服务 + Hertz HTTP服务完成的第五届字节跳动青训营-极简抖音后端项目

## 一、项目特点

1. 采用RPC框架（Kitex）和 Hertz 脚手架生成代码进行开发，基于 **RPC 微服务** + **Hertz 提供 HTTP 服务**

2. 基于《[接口文档在线分享](https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/)[- Apifox](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)》提供的接口进行开发，使用《[极简抖音](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7)[App使用说明 - 青训营版](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 》提供的APK进行Demo测试， **功能完整实现** ，前端接口匹配良好。

3. 代码结构采用 (HTTP API 层 + RPC Service 层+Dal 层) 项目 **结构清晰** ，代码 **符合规范**

4. 使用 **JWT** 进行用户token的校验

5. 使用 **ETCD** 进行服务发现和服务注册；

6. 使用 **Minio** 实现视频文件和图片的对象存储

7. 使用 **Gorm、Gorm-gen** 对 MySQL 进行 ORM 操作；

8. 使用 **Redis** 实现点赞和社交业务 Dal 层及部分缓存；

## 二、项目地址

- **<https://github.com/ClubWeGo/simple-tiktok>**
- **<http://124.221.147.131:8888>**

## 三、项目说明

### 1. 项目模块介绍
| 服务名称                | 模块介绍              | 技术框架                      | 传输协议  | 注册中心  | 数据存储   |
|---------------------|-------------------|---------------------------|-------|-------|--------| 
| biz                 | HTTP请求处理，RPC调用    | `JWT` `Kitex` `Hertz`     | `http` | `etcd` | `Minio` |
| services/usermicro  | 用户微服务             | `Gorm` `Kitex` `Hertz`    | `thrift` | `etcd` | `MySQL` `gorm` |
| services/relationmicro | 社交微服务             | `Redis` `Kitex`           | `thrift` | `etcd` | `Redis` |
| services/videomicro | 视频微服务             | `Gorm` `Kitex` `Minio`    | `thrift` | `etcd` | `MySQL` `Minio` |
| services/favoritemicro | 点赞微服务             | `Redis`  `Kitex`          | `thrift` | `etcd` | `Redis` |
| services/commentmicro | 评论微服务             | `gorm` `Kitex`            | `thrift` | `etcd` | `MySQL` `gorm` |
| services/*/dal      | 数据层实现             | `MySQL` `gorm` `gorm-gen` | -     | -     | `MySQL` `gorm` |  

### 2. 服务调用关系

- HTTP 使用 Hertz 开放 HTTP 端口, 提供鉴权和部分错误处理, 通过封装的RPC客户端与微服务中的服务端通信;

- RPC 微服务, 接收客户端的请求, 在各自的 handler 中实现与数据库交互的业务逻辑;

- DAL 提供数据层实现, pack 部分实现将数据库输出封装为服务端的响应结构体;



## 四、Docker 部署
1. 拉取代码
    ```bash
    git clone https://github.com/ClubWeGo/simple-tiktok.git
    ```
2. 打开目录
    ```bash
    cd simple-tiktok
    ```
3. 构建镜像
    ```bash
    docker build -t usermicro:latest ./services/usermicro
    docker build -t videomicro:latest ./services/videomicro
    docker build -t favoritemicro:latest ./services/favoritemicro
    docker build -t commentmicro:latest ./services/commentmicro
    docker build -t relationmicro:latest ./services/relationmicro
    docker build -t simple-tiktok:latest .
    ```
4. compose 启动容器
    ```bash
    docker-compose up -d
   
5. 访问
    http://localhost:8888/ping

## 五、手动部署
### 生成代码
`hz new -idl idl/core.thrift -module github.com/ClubWeGo/douyin`

### 扩展模块
> `hz new -idl idl/interaction.thrift -module github.com/ClubWeGo/douyin`
> `hz new -idl idl/relation.thrift -module github.com/ClubWeGo/douyin`

### 删除代码
`rm -r ./biz router.go router_gen.go main.go go.mod go.sum .hz .gitignore`

### 更新代码
`hz update -idl idl/core.thrift`

### 拉取依赖
`go mod tidy`

### 编译代码
`go build`


依赖另外几个微服务的代码仓库

(微服务代码更新后，需主动执行来更新代码)

go get github.com/ClubWeGo/simple-tiktok/services/videomicro@latest

go get github.com/ClubWeGo/simple-tiktok/services/usermicro@latest

go get github.com/ClubWeGo/simple-tiktok/services/relationmicro@latest

go get github.com/ClubWeGo/simple-tiktok/services/favoritemicro@latest

go get github.com/ClubWeGo/simple-tiktok/services/commentmicro@latest


# 说明
1. 当前注册用户时，前端并未提供个人参数配置的能力，也没有提供更新用户信息的能力，所以在core.register_server处写死了初始化的用户背景图像和头像。在配置本项目时，请将对应文件名换成minio中存储的文件名。
