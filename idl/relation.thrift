namespace go interaction

struct User {
    1: required i64 id;
    2: required string name;
    3: optional i64 follow_count;
    4: optional i64 follower_count;
    5: required bool is_follow;
}


# followList
struct FollowListReq {
    1: required string token;
    2: required i64 user_id;
}

struct FollowListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: list<User> user_list;
}

service FollowListService {
    FollowListResp FollowListMethod(1: FollowListReq request) (api.get="/douyin/relation/follow/list/")
}


# followerList
struct FollowerListReq {
    1: required string token;
    2: required i64 user_id;
}

struct FollowerListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: list<User> user_list;
}

service FollowerListService {
    FollowerListResp FollowerListMethod(1: FollowerListReq request) (api.get="/douyin/relation/follower/list/")
}


# friendList
struct FriendListReq {
    1: required string token;
    2: required i64 user_id;
}

struct FriendListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: list<User> user_list;
}

service FriendListService {
    FriendListResp FriendListMethod(1: FriendListReq request) (api.get="/douyin/relation/friend/list/")
}


# relation
struct RelationReq {
    1: required string token;
    2: required i64 to_user_id;
    3: required i32 action_type; // 1 subscribe; 2 cancle
}

struct RelationResp {
    1: required i32 status_code;
    2: optional string status_msg;
}

service RelationService {
    RelationResp RelationMethod(1: RelationReq request) (api.post="/douyin/relation/action/")
}