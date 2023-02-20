namespace go favorite


struct BaseResp {
    1: required i32 status_code;
    2: optional string status_msg;
}


# favoriteList
struct FavoriteListReq {
    1: required i64 user_id;
}

struct FavoriteListResp {
    1: required BaseResp base_resp
    2: required list<i64> video_id_list;
}


# favorite
struct FavoriteReq {
    1: required i64 user_id;
    2: required i64 video_id;
    3: required i32 action_type;
}

struct FavoriteResp {
    1: required BaseResp base_resp;
}


struct FavoriteRelationReq {
    1: required i64 user_id;
    2: required i64 video_id;
}

struct FavoriteRelationResp {
    1: required BaseResp base_resp;
    2: required bool is_favorite;
}


struct VideoFavoriteCountReq {
    1: required i64 video_id;
}

struct VideoFavoriteCountResp {
    1: required BaseResp base_resp;
    2: required i64 favorite_count;  // 被点赞数量
}


struct UserFavoriteCountReq {
    1: required i64 user_id;
}

struct UserFavoriteCountResp {
    1: required BaseResp base_resp
    2: required i64 favorite_count;  // 点赞数量
    3: required i64 favorited_count;  // 被点赞数量
}


struct VideosFavoriteCountReq {
    1: required list<i64> video_id_list;
}

struct VideosFavoriteCountResp {
    1: required BaseResp base_resp;
    2: required map<i64, i64> favorite_count_map;
}


struct UsersFavoriteCountReq {
    1: required list<i64> user_id_list;
}

struct UsersFavoriteCountResp {
    1: required BaseResp base_resp;
    3: required map<i64, list<i64>> favorite_count_map;
}


service FavoriteService {
    FavoriteResp FavoriteMethod(1: FavoriteReq request);
    FavoriteListResp FavoriteListMethod(1: FavoriteListReq request)
    FavoriteRelationResp FavoriteRelationMethod(1: FavoriteRelationReq request);
    VideoFavoriteCountResp VideoFavoriteCountMethod(1: VideoFavoriteCountReq request);
    UserFavoriteCountResp UserFavoriteCountMethod(1: UserFavoriteCountReq request);
    VideosFavoriteCountResp VideosFavoriteCountMethod(1: VideosFavoriteCountReq request);
    UsersFavoriteCountResp UsersFavoriteCountMethod(1: UsersFavoriteCountReq request);
}
