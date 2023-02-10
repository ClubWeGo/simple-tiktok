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
	data = append(data, Video{
		"Introducing Chromecast. The easiest way to enjoy online video", 2, "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerEscapes.mp4", "https://i.imgtg.com/2022/05/06/zJFbg.th.jpg",
	})
	data = append(data, Video{
		"http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerBlazes.mp4", 3, "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerFun.mp4", "https://i.imgtg.com/2022/05/06/zJAiC.th.jpg",
	})
	data = append(data, Video{
		"Introducing Chromecast. The easiest way to enjoy onlin", 1, "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerJoyrides.mp4", "https://i.imgtg.com/2022/05/06/zJK4L.th.png",
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
