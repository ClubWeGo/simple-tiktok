namespace go usermicro

// 所有字段的约束和检查交给业务端处理（如字段的异常检测，非法值，黑灰产处理等），微服务只负责功能逻辑
// 虽然userInfo也存了一些点赞数之类的数据，但最新数据还是从对应微服务获取，这里只是缓存
struct UserInfo {
    1: required i64 id;
    2: required string name;
    3: optional string email; //(vt.pattern = "^[A-Za-z0-9-_\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$"); // https://www.jianshu.com/p/5966a2d9df75
                            // bug，正则验证邮箱无法通过
    4: required i64 follow_count; // 关注总数
    5: required i64 follower_count; // 粉丝总数
    6: required string avatar; // 头像地址
    7: required string background_image; // 背景图地址
    8: required string signature; // 个人简介
    9: required i64 total_favorited; // 获赞数量
    10: required i64 work_count; // 作品数
    11: required i64 favorite_count; // 喜欢数
}

struct CreateUserInfo {
    1: required string name;
    2: optional string email; //(vt.pattern = "^[A-Za-z0-9-_\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$"); // https://www.jianshu.com/p/5966a2d9df75
                            // bug，正则验证邮箱无法通过
    3: optional string avatar; // 头像地址
    4: optional string background_image; // 背景图地址
    5: optional string signature; // 个人简介
}

// 可被更新的用户信息字段
struct UpdateUserInfo {
    1: optional string email;
    2: optional string password;
    3: optional string avatar; // 头像地址
    4: optional string background_image; // 背景图地址
    5: optional string signature; // 个人简介
}

// 更新社交信息缓存
struct UpdateRelationCache {
    1: required i64 follow_count; // 关注总数
    2: required i64 follower_count; // 粉丝总数
}

// 更新交互信息缓存
struct UpdateInteractionCache {
    1: required i64 total_favorited; // 获赞数量
    2: required i64 favorite_count; // 喜欢数
}

// 更新视频数量缓存
struct UpdateWorkCache {
    1: required i64 work_count; // 作品数
}

// 用户登录校验，支持name和email
struct LoginUserReq {
    1: required string name;
    3: required string password;
}

struct LoginUserResp {
    1: required bool status;
    // 如果成功返回用户信息, 目前app只需要user_id
    2: optional i64 user_id;
}

// 根据id或者name查询用户
struct GetUserReq {
    1: optional i64 id; // 数据库只在id和name设置索引，尽量只用id和name查询
    2: optional string name;
}

struct GetUserResp {
    1: required bool status;
    // 如果成功返回用户信息
    2: optional UserInfo user;
}

// 批量查询用户
struct GetUserSetByIdSetReq {
    1: optional list<i64> id_set;  // 批量的id查询批量的用户
}

struct GetUserSetByIdSetResp {
    1: required bool status;
    2: optional list<UserInfo> user_set;
}

// 创建用户：目前app仅要求使用用户名和密码创建用户
// 但是考虑到后续，用户多样化信息注册实习，这个逻辑应该交给app，所以采用对象来注册
struct CreateUserReq {
    1: required CreateUserInfo newuser;
    2: required string password; // 微服务进行md5操作
}

struct CreateUserResp {
    1: required bool status;
    // 如果成功返回用户信息，app仅需要user_id
    2: optional i64 user_id;
}

// 更新用户信息，更新是用户私人操作，肯定有id信息，要求使用id进行请求
struct UpdateUserReq {
    1: required i64 id; // 用id去更新条目
    3: required UpdateUserInfo update_data;
}

struct UpdateUserResp {
    1: required bool status;
}

// 更新缓存的接口，由业务代码控制更新策略
struct UpdateRelationCacheReq {
    1: required i64 id;
    2: required UpdateRelationCache new_data;
}

struct UpdateRelationCacheResp {
    1: required bool status;
}

struct UpdateInteractionCacheReq {
    1: required i64 id;
    2: required UpdateInteractionCache new_data;
}

struct UpdateInteractionCacheResp {
    1: required bool status;
}

struct UpdateWorkCacheReq {
    1: required i64 id;
    2: required UpdateWorkCache new_data;
}

struct UpdateWorkCacheResp {
    1: required bool status;
}

service UserService {
    GetUserResp GetUserMethod(1: GetUserReq request)
    GetUserSetByIdSetResp GetUserSetByIdSetMethod(1: GetUserSetByIdSetReq request)
    LoginUserResp LoginUserMethod(1: LoginUserReq request)
    CreateUserResp CreateUserMethod(1: CreateUserReq request)
    UpdateUserResp UpdateUserMethod(1: UpdateUserReq request)

    // 缓存更新接口
    UpdateRelationCacheResp UpdateRelationMethod(1: UpdateRelationCacheReq request)
    UpdateInteractionCacheResp UpdateInteractionMethod( 1: UpdateInteractionCacheReq request)
    UpdateWorkCacheResp UpdateWorkMethod(1: UpdateWorkCacheReq request)
}