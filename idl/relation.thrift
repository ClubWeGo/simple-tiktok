namespace go relation

struct User {
    1:  required i64        id;
    2:  required string     name;               // 用户昵称
    3:  optional i64        follow_count;       // 关注数
    4:  optional i64        follower_count;     // 粉丝数
    5:  required bool       is_follow;          // 是否关注 true-已关注 false-未关注
    6:  optional string     avatar;             // 头像
    7:  optional string     background_image;   // 个人顶部大图
    8:  optional string     signature;          // 个人简介
    9:  optional i64        total_favorited;    // 获赞数
    10: optional i64        work_count;         // 作品数
    11: optional i64        favorite_count;     // 喜欢数
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


struct Message {
    1: required i64 id; // 消息id
    2: required i64 to_user_id; // 该消息接收者的id
    3: optional i64 from_user_id; // 该消息发送者的id
    4: optional string content; // 消息内容
    5: required string create_time; // 消息创建时间
}

# Get P2P Messages
struct MessageChatReq {
    1: required string token; 
    2: required i64 to_user_id;
}

struct MessageChatResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required list<Message> message_list;
}

service MessageChatService {
    MessageChatResp MessageChatMethod(1: MessageChatReq request) (api.post="/douyin/message/chat/")
}

# Send Messages
struct MessageActionReq {
    1: required string token; 
    2: required i64 to_user_id;
    3: required i32 action_type; // 1 表示发送消息
    4: required string content; // 消息内容
}

struct MessageActionResp {
    1: required i32 status_code;
    2: optional string status_msg;
}

service MessageActionService {
    MessageActionResp MessageActionMethod(1: MessageActionReq request) (api.post="/douyin/message/action/")
}
=======
namespace go relation
include "core.thrift"



# followList
struct FollowListReq {
    1: required string token;
    2: required i64 user_id;
}

struct FollowListResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: list<core.User> user_list;
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
    3: list<core.User> user_list;
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
    3: list<core.User> user_list;
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


struct Message {
    1: required i64 id; // 消息id
    2: required i64 to_user_id; // 该消息接收者的id
    3: optional i64 from_user_id; // 该消息发送者的id
    4: optional string content; // 消息内容
    5: required string create_time; // 消息创建时间
}

# Get P2P Messages
struct MessageChatReq {
    1: required string token;
    2: required i64 to_user_id;
}

struct MessageChatResp {
    1: required i32 status_code;
    2: optional string status_msg;
    3: required list<Message> message_list;
}

service MessageChatService {
    MessageChatResp MessageChatMethod(1: MessageChatReq request) (api.post="/douyin/message/chat/")
}

# Send Messages
struct MessageActionReq {
    1: required string token;
    2: required i64 to_user_id;
    3: required i32 action_type; // 1 表示发送消息
    4: required string content; // 消息内容
}

struct MessageActionResp {
    1: required i32 status_code;
    2: optional string status_msg;
}

service MessageActionService {
    MessageActionResp MessageActionMethod(1: MessageActionReq request) (api.post="/douyin/message/action/")
}
>>>>>>> 8563607bdd664d6d1ed1bb1a786b655ea52aba58
