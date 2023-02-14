namespace go favorite

struct User {
    1: required i64 id;
    2: required string name;
    3: optional i64 follow_count;
    4: optional i64 follower_count;
    5: required bool is_follow;
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


# favoriteList
struct FavoriteListReq {
    1: required string token;
    2: required i64 user_id;
}

struct FavoriteListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required list<Video> video_list;
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


service FavoriteService {
    FavoriteResp FavoriteMethod(1: FavoriteReq request);
    FavoriteListResp FavoriteListMethod(1: FavoriteListReq request)
}
