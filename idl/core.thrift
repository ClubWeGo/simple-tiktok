namespace go core

struct User {
    1: required i64 id;
    2: required string name;
    3: required i64 follow_count;
    4: required i64 follower_count;
    5: required bool is_follow;
    6: required string avatar;
    7: required string background_image;
    8: required string signature;
    9: required string total_favourited;
    10: required i64 work_count;
    11: required i64 favorite_count;
}

struct Video {
    1: required i64 id;
    2: required User author;
    3: required string play_url;
    4: required string cover_url;
    5: required i64 favorite_count;
    6: required i64 comment_count;
    7: required bool is_favorite;
    8: required string title;
}


# Feed
struct FeedReq {
    1: optional i64 latest_time; // 请求时间点
    2: optional string token;
}

struct FeedResp {
    1: required i32 status_code;  // 0 success; 1 failed
    2: optional string status_msg;
    3: required list<Video> video_list;
    4: required i64 next_time; // 本次请求到数据中最早的时间作为下次请求的开始时间
}

service FeedServer {
    FeedResp FeedMethod(1: FeedReq request) (api.get="/douyin/feed/");
}


# login
struct LoginReq {
    1: required string username (api.vd = "(len($) > 0 && len($) < 33); msg:sprintf('invalid account_type: %v',$)"); // 文档要求32个字符
    2: required string password (api.vd = "(len($) > 0 && len($) < 33); msg:sprintf('invalid account_type: %v',$)");
}

struct LoginResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required i64 user_id;
    4: required string token;
}

service LoginServer {
    LoginResp LoginMethod(1: LoginReq request) (api.post="/douyin/user/login/");
}


# publishAction
struct PublishActionReq {
    1: required string token;
    2: required binary data;
    3: required string title; // 虽然没有写，但是标题长度也是得做限制的；需要考虑videomicro的表设计
}

struct PublishActionResp {
    1: required i32 status_code;
    2: optional string status_msg;
}

service PublishActionServer {
    PublishActionResp PublishActionMethod(1: PublishActionReq request) (api.post="/douyin/publish/action/");
}


# publishList
struct PublishListReq {
    1: required i64 user_id (api.vd = "$>0; msg:sprintf('invalid user_id: %v',$)");
    2: required string token;
}

struct PublishListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required list<Video> video_list;
}

service PublishListServer {
    PublishListResp PublishListMethod(1: PublishListReq request) (api.get="/douyin/publish/list/");
}


# register
struct RegisterReq {
    1: required string username (api.vd = "(len($) > 0 && len($) < 33); msg:sprintf('invalid account_type: %v',$)");
    2: optional string password (api.vd = "(len($) > 0 && len($) < 33); msg:sprintf('invalid account_type: %v',$)");
}

struct RegisterResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required i64 user_id;
    4: required string token;
}

service RegisterServer {
    RegisterResp RegisterMethod(1: RegisterReq request) (api.post="/douyin/user/register/");
}


# userinfo
struct UserInfoReq {
    1: required i64 user_id (api.vd = "$>0; msg:sprintf('invalid user_id: %v',$)");
    2: required string token;
}

struct UserInfoResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required User user;
}

service UserInfoServer {
    UserInfoResp UserInfoMethod(1: UserInfoReq request) (api.get="/douyin/user/");
}