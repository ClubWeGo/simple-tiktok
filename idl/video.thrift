namespace go videomicro

struct Video {
    1: required i64 id;
    2: required string title (vt.min_size = "6", vt.max_size = "126");
    3: required i64 author_id;
    4: required string play_url;
    5: required string cover_url;
    6: required i64 favorite_count;
    7: required i64 comment_count;
}

struct CreateVideoReq {
    1: required string title (vt.min_size = "6", vt.max_size = "126");
    2: required i64 author_id;
    3: required string play_url;
    4: required string cover_url;
}

struct CreateVideoResp {
    1: required bool status;
}

// get video by video id
struct GetVideoReq {
    1: required i64 id;
}

struct GetVideoResp {
    1: required bool status;
    2: optional Video video;
}

// get video list by author id
struct GetVideosByAuthorIdReq {
    1: required i64 author_id;
}

struct GetVideosByAuthorIdResp {
    1: required bool status;
    2: optional list<Video> video_list; 
}

// get video feed list
struct GetVideosFeedReq {
    1: required i64 latest_time; // hertz端如果有next time则用于latest_time，否则为最新时间的time.Now().Unix()
    2: required i32 limit;
}
// 返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
struct GetVideosFeedResp {
    1: required bool status;
    2: optional i64 next_time;
    3: optional list<Video> video_list;
}

// delete video; 鉴权交予业务层处理
struct UpdateVideoReq {
    1: required i64 id;
    2: optional string title; // 当前简化版本，只有title是可以由用户修改的
}

struct UpdateVideoResp {
    1: required bool status;
    2: optional Video video;
}

// delete video; 鉴权交予业务层处理
struct DeleteVideoReq {
    1: required i64 id
}

struct DeleteVideoResp {
    1: required bool status;
}

service VideoService {
    CreateVideoResp CreateVideoMethod(1: CreateVideoReq request)
    GetVideoResp GetVideoMethod(1: GetVideoReq request)
    GetVideosByAuthorIdResp GetVideosByAuthorIdMethod(1: GetVideosByAuthorIdReq request)
    GetVideosFeedResp GetVideosFeedMethod(1: GetVideosFeedReq request)
    UpdateVideoResp UpdateVideoMethod(1: UpdateVideoReq request)
    DeleteVideoResp DeleteVideoMethod(1: DeleteVideoReq request)
}

