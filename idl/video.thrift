namespace go videomicro

struct Video {
    1: required i64 id;
    2: required string title;
    3: required i64 author_id;
    4: required string play_url;
    5: required string cover_url;
    6: required i64 favorite_count; // 视频点赞数缓存
    7: required i64 comment_count; // 评论数缓存 默认从交互微服务拿数据，这里仅作缓存
}

// 创建视频
struct CreateVideoReq {
    1: required string title;
    2: required i64 author_id;
    3: required string play_url;
    4: required string cover_url;
}

struct CreateVideoResp {
    1: required bool status;
}

// videoid查video全部信息，使用GetVideoSetByIdSetReq

// 通过 videoid 查 author id
struct GetVideoAuthorIdReq {
    1: required i64 id;
}

struct GetVideoAuthorIdResp {
    1: required bool status;
    2: required i64 author_id;
}

// 用户已发布的视频
struct GetVideosByAuthorIdReq {
    1: required i64 author_id;
}

struct GetVideosByAuthorIdResp {
    1: required bool status;
    2: optional list<Video> video_list; 
}

// feed流
struct GetVideosFeedReq {
    1: required i64 latest_time; // hertz端如果有next time则用于latest_time，否则为最新时间的time.Now().Unix()
    // 统一使用19位的纳秒时间戳，为满足不同应用要求
    2: required i32 limit;
}
// 返回按投稿时间倒序的视频列表，视频数由应用控制传入limit。
struct GetVideosFeedResp {
    1: required bool status;
    // 统一使用19位的纳秒时间戳，为满足不同应用要求
    2: optional i64 next_time;
    3: optional list<Video> video_list;
}

// 更新视频; 鉴权交予业务层处理
struct UpdateVideoReq {
    1: required i64 id;
    2: optional string title; // 当前简化版本，只有title是可以由用户修改的
}

struct UpdateVideoResp {
    1: required bool status;
}

// 删除; 鉴权交予业务层处理
struct DeleteVideoReq {
    1: required i64 video_id;
}

struct DeleteVideoResp {
    1: required bool status;
}

struct VideoCount {
    1: required i64 id;
    2: required i64 count;
}

// 统计用户的发布视频数量
struct GetVideoCountSetByIdUserSetReq {
    1: required list<i64> author_id_set;
}

struct GetVideoCountSetByIdUserSetResp {
    1: required bool status;
    2: required list<VideoCount> count_set;
}

// 批量查询视频；传入视频id切片，返回视频对象切片;如果查询单挑视频，传入一个id
struct GetVideoSetByIdSetReq {
    1: required list<i64> id_set;  // 批量的video id查询批量的视频
}

struct GetVideoSetByIdSetResp {
    1: required bool status;
    2: optional list<Video> video_set;
}

service VideoService {
    CreateVideoResp CreateVideoMethod(1: CreateVideoReq request)
    GetVideoAuthorIdResp GetVideoAuthorIdMethod(1: GetVideoAuthorIdReq request)
    GetVideosByAuthorIdResp GetVideosByAuthorIdMethod(1: GetVideosByAuthorIdReq request)
    GetVideoSetByIdSetResp GetVideoSetByIdSetMethod(1: GetVideoSetByIdSetReq request)
    GetVideosFeedResp GetVideosFeedMethod(1: GetVideosFeedReq request)
    UpdateVideoResp UpdateVideoMethod(1: UpdateVideoReq request)
    DeleteVideoResp DeleteVideoMethod(1: DeleteVideoReq request)

    // 获取用户的作品数量接口
    GetVideoCountSetByIdUserSetResp GetVideoCountSetByIdUserSetMethod(1: GetVideoCountSetByIdUserSetReq request)
}

