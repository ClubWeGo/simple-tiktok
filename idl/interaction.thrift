namespace go interaction
include "core.thrift"



struct Comment {
    1: required i64 id;
    2: required core.User user;
    3: required string content;
    4: required string create_date;
}


# commentList
struct CommentListReq {
    1: required string token;
    2: required i64 video_id;
}

struct CommentListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required list<Comment> comment_list;
}

service CommentListServer {
    CommentListResp CommentListMethod(1: CommentListReq request) (api.get="/douyin/comment/list/");
}


# comment
struct CommentReq {
    1: required string token;
    2: required i64 video_id;
    3: required i32 action_type; // 1 publish, 2 delite
    4: optional string comment_text;
    5: optional i64 comment_id;
}

struct CommentResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: optional Comment comment;
}

service CommentServer {
    CommentResp CommentMethod(1: CommentReq request) (api.post="/douyin/comment/action/");
}


# favoriteList
struct FavoriteListReq {
    1: required string token;
    2: required i64 user_id;
}

struct FavoriteListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required list<core.Video> video_list;
}

service FavoriteListServer {
    FavoriteListResp FavoriteListMethod(1: FavoriteListReq request) (api.get="/douyin/favorite/list/");
}


# favorite
struct FavoriteReq {
    1: required string token;
    2: required i64 video_id;
    3: required i32 action_type;
}

struct FavoriteResp {
    1: required i32 status_code;
    2: optional string status_msg;
}

service FavoriteServer {
    FavoriteResp FavoriteMethod(1: FavoriteReq request) (api.post="/douyin/favorite/action/");
}
