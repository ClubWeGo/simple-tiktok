package main

import (
	"context"
	"log"
	"time"

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
	// Thanks https://www.jianshu.com/p/d5f00ad58572

	data = append(data, Video{
		"test1", 1, "https://media.w3.org/2010/05/sintel/trailer.mp4", "https://i.imgtg.com/2022/05/06/zJAiC.th.jpg",
	})
	data = append(data, Video{
		"test2", 2, "http://www.w3school.com.cn/example/html5/mov_bbb.mp4", "https://i.imgtg.com/2022/05/06/zJiVs.th.jpg",
	})
	data = append(data, Video{
		"test3", 1, "https://www.w3schools.com/html/movie.mp4", "https://i.imgtg.com/2022/05/06/zJNQS.th.jpg",
	})
	data = append(data, Video{
		"test4", 2, "http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4", "https://i.imgtg.com/2022/05/06/zJFbg.th.jpg",
	})
	data = append(data, Video{
		"test5", 3, "https://player.vimeo.com/external/188350983.sd.mp4?s=0bdf01fb5f5c66e43ddae76f573cef2a7786de64&profile_id=164", "https://i.imgtg.com/2022/05/06/zJAiC.th.jpg",
	})
	data = append(data, Video{
		"test6", 1, "https://player.vimeo.com/external/188355959.sd.mp4?s=e5eea0f749282013db81a7e5cd047c57e066e2b9&profile_id=164", "https://i.imgtg.com/2022/05/06/zJK4L.th.png",
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

	// create video
	testdata := generateTestData()
	for _, video := range testdata {
		resp, err := client.CreateVideoMethod(context.Background(), &videomicro.CreateVideoReq{
			Title:    video.Title,
			AuthorId: video.Author_id,
			PlayUrl:  video.Play_url,
			CoverUrl: video.Cover_url,
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second * 2)
	}

	// // get video
	// resp1, err := client.GetVideoMethod(context.Background(), &videomicro.GetVideoReq{Id: 23})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp1)

	// // getFeed
	// resp, err := client.GetVideosFeedMethod(context.Background(), &videomicro.GetVideosFeedReq{Offset: 0, Limit: 30})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp)

	// // getUser's Videos
	// resp, err := client.GetVideosByAuthorIdMethod(context.Background(), &videomicro.GetVideosByAuthorIdReq{AuthorId: 3, Offset: 0, Limit: 30})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp)

	// // // update video
	// newtitle := "newtitlefor23"
	// resp2, err := client.UpdateVideoMethod(context.Background(), &videomicro.UpdateVideoReq{Id: 23, Title: &newtitle})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp2)

	// // delete video
	// resp, err := client.DeleteVideoMethod(context.Background(), &videomicro.DeleteVideoReq{Id: 23})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp)
}
