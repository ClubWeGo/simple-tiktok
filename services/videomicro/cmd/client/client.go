package main

import (
	"context"
	"log"

	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro/videoservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type Video struct {
	Title     string
	Author_id int64
	Play_url  string
	Cover_url string
}

func generateTestData() []Video {
	data := make([]Video, 0)
	// Thanks https://gist.github.com/jsturgis/3b19447b304616f18657

	data = append(data, Video{
		"Big Buck Bunny tells the story of a giant rabbit", 1, "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4", "https://i.imgtg.com/2022/05/06/zJAiC.th.jpg",
	})
	data = append(data, Video{
		"The first Blender Open Movie from 2006", 2, "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4", "https://i.imgtg.com/2022/05/06/zJiVs.th.jpg",
	})
	data = append(data, Video{
		"HBO GO now works with Chromecast", 1, "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerBlazes.mp4", "https://i.imgtg.com/2022/05/06/zJNQS.th.jpg",
	})
	return data
}

func main() {
	r, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	client, err := videoservice.NewClient("videoservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	// // create video
	// testdata := generateTestData()
	// for _, video := range testdata[1:] {
	// 	resp, err := client.CreateVideoMethod(context.Background(), &videomicro.CreateVideoReq{
	// 		Title:    video.Title,
	// 		AuthorId: video.Author_id,
	// 		PlayUrl:  video.Play_url,
	// 		CoverUrl: video.Cover_url,
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(resp)
	// 	time.Sleep(time.Second * 2)
	// }

	// // get video
	// resp1, err := client.GetVideoSetByIdSetMethod(context.Background(), &videomicro.GetVideoSetByIdSetReq{
	// 	IdSet: []int64{1, 2, 5}, // 对于不存在的id，会没有这一项内容
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp1)

	// // getFeed
	// latestTime := time.Now().UnixNano() //与rpc通信 统一使用纳秒
	// for i := 0; i < 1; i++ {
	// 	resp, err := client.GetVideosFeedMethod(context.Background(), &videomicro.GetVideosFeedReq{LatestTime: latestTime, Limit: 5})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	latestTime = *resp.NextTime
	// 	log.Println(resp)
	// }

	// // getUser's Videos
	// resp, err := client.GetVideosByAuthorIdMethod(context.Background(), &videomicro.GetVideosByAuthorIdReq{AuthorId: 1})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp)

	// // // update video
	// newtitle := "newtitlefor23"
	// resp2, err := client.UpdateVideoMethod(context.Background(), &videomicro.UpdateVideoReq{Id: 11, Title: &newtitle})
	// if err != nil { // 如果传一个错误id进去，返回也是成功，但是不存在修改
	// 	log.Fatal(err)
	// }
	// log.Println(resp2)

	// // delete video
	// resp, err := client.DeleteVideoMethod(context.Background(), &videomicro.DeleteVideoReq{VideoId: 23})
	// if err != nil { // 删除一个不存在的id，record not found
	// 	log.Fatal(err)
	// }
	// log.Println(resp)

	// 获取用户发布的视频数
	resp, err := client.GetVideoCountSetByIdUserSetMethod(context.Background(), &videomicro.GetVideoCountSetByIdUserSetReq{
		AuthorIdSet: []int64{1, 2},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)

	// // 获取视频对应作者id
	// resp, err := client.GetVideoAuthorIdMethod(context.Background(), &videomicro.GetVideoAuthorIdReq{
	// 	Id: 12,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp)

}
