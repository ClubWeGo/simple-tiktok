package main

import (
	"github.com/ClubWeGo/simple-tiktok/services/commentmicro/dal"
	comment "github.com/ClubWeGo/simple-tiktok/services/commentmicro/kitex_gen/comment/commentservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {

	dal.Init()

	r, err := etcd.NewEtcdRegistry([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:10010")
	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "commentservice"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
