package main

import (
	"context"
	"github.com/ClubWeGo/commentmicro/kitex_gen/comment"
	"github.com/ClubWeGo/commentmicro/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"time"
)

func main() {

	r, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}
	c, err := commentservice.NewClient("commentservice",
		client.WithResolver(r),
		//client.WithHostPorts("0.0.0.0:10010"),
		client.WithMuxConnection(1),
	)
	if err != nil {
		log.Fatal(err)
	}
	//commenttext := "hhh"
	//req := &comment.CommentReq{UserId: 1, VideoId: 1, ActionType: 1, CommentText: &commenttext}
	//resp, err := c.CommentMethod(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)
	//
	//commenttext := "hi"
	//id := int64(2)
	//req := &comment.CommentReq{ActionType: 2, CommentId: &id, CommentText: &commenttext}
	//resp, err := c.CommentMethod(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)

	req := &comment.CommentListReq{UserId: 1, VideoId: 1}
	resp, err := c.CommentListMethod(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	//var sliceName []int64 = []int64{1, 2}
	//req := &comment.VideosCommentCountReq{sliceName}
	//resp, err := c.VideosFavoriteCountMethod(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(resp)
}
