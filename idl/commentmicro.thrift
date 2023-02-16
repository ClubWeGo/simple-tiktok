namespace go comment


struct User {
    1: required i64 id;
    2: required string name;
    3: optional i64 follow_count;//关注总数
    4: optional i64 follower_count;//粉丝总数
    5: required bool is_follow;//true 已关注 false未关注该视频主
}

struct Comment {
    1: required i64 id; //视频评论id
    2: required User user;//评论用户的id
    3: required string content; //评论内容
    4: required string create_date;//评论发布日期时间 mm-dd
}


# commentList
struct CommentListReq {
    1: required string token;//用户鉴权token
    2: required i64 video_id;//视频id
}

struct CommentListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required list<Comment> comment_list;
}


# comment
struct CommentReq {
//    1: required string token;
    1: required i64 user_id;
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


# service
service CommentService {
    CommentResp CommentMethod(1: CommentReq request) ;
    CommentListResp CommentListMethod(1: CommentListReq request)
}
