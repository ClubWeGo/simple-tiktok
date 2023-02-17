namespace go favorite


# favoriteList
struct FavoriteListReq {
    1: required i64 user_id;
}

struct FavoriteListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required list<i64> video_id_list;
}


# favorite
struct FavoriteReq {
    1: required i64 user_id;
    2: required i64 video_id;
    3: required i32 action_type;
}

struct FavoriteResp {
    1: required i32 status_code;
    2: optional string status_msg;
}


struct FavoriteRelationReq {
    1: required i64 user_id;
    2: required i64 video_id;
}

struct FavoriteRelationResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required bool is_favorite;
}


struct VideoFavoriteCountReq {
    1: required i64 video_id;
}

struct VideoFavoriteCountResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required i64 favorite_count;
}


struct UserFavoriteCountReq {
    1: required i32 user_id;
}

struct UserFavoriteCountResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required i64 favorite_count;
    4: required i64 favorited_count;
}


service FavoriteService {
    FavoriteResp FavoriteMethod(1: FavoriteReq request);
    FavoriteListResp FavoriteListMethod(1: FavoriteListReq request)
    FavoriteRelationResp FavoriteRelationMethod(1: FavoriteRelationReq request);
    VideoFavoriteCountResp VideoFavoriteCountMethod(1: VideoFavoriteCountReq request);
    UserFavoriteCountResp UserFavoriteCountMethod(1: UserFavoriteCountReq request);
}
