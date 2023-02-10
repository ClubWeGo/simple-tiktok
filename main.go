package main

import (
	"log"
	"net"

	"github.com/ClubWeGo/videomicro/dal"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro/videoservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	dsn := "tk:123456@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
	dal.InitDB(dsn)

	r, err := etcd.NewEtcdRegistry([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:10001")
	svr := videoservice.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "videoservice"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
