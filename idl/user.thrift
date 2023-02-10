namespace go usermicro

struct UserInfo {
    1: required i64 id;
    2: required string name (vt.min_size = "6", vt.max_size = "30"); // 汉字长度为3，此处限制2~10个汉字
    3: optional string email; //(vt.pattern = "^[A-Za-z0-9-_\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$"); // https://www.jianshu.com/p/5966a2d9df75
                            // bug，正则验证邮箱无法通过
    4: required i64 follow_count;
    5: required i64 follower_count;
}

// login User
struct LoginUserReq {
    1: optional string name (vt.min_size = "6", vt.max_size = "30");
    2: optional string email;
    3: required string password;
}

struct LoginUserResp {
    1: required bool status;
    2: optional UserInfo user; // 如果成功返回用户信息
}

// get User
struct GetUserReq {
    1: optional i64 id; // 数据库只在id和name设置索引，尽量只用id和name查询
    2: optional string name (vt.min_size = "6", vt.max_size = "30");
    3: optional string email;
}

struct GetUserResp {
    1: required bool status;
    2: optional UserInfo user; // 如果成功返回用户信息
}

// CreateUser
struct CreateUserReq {
    1: required string name (vt.min_size = "6", vt.max_size = "30");
    2: required string password; // 微服务进行md5操作
    3: optional string email;
}

struct CreateUserResp {
    1: required bool status;
    2: optional UserInfo user; // 如果成功返回用户信息
}

// UpdateUser
struct UpdateUserReq {
    1: optional i64 id; // 用id去更新条目 索引
    2: optional string name (vt.min_size = "6", vt.max_size = "30"); // 用name更新条目 索引     
    3: optional string email;
    4: optional string password;
}

struct UpdateUserResp {
    1: required bool status;
    2: optional UserInfo user; // 如果成功返回用户信息
}

service UserService {
    GetUserResp GetUserMethod(1: GetUserReq request)
    LoginUserResp LoginUserMethod(1: LoginUserReq request)
    CreateUserResp CreateUserMethod(1: CreateUserReq request)
    UpdateUserResp UpdateUserMethod(1: UpdateUserReq request)
}