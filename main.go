package main

import (
	"github.com/ClubWeGo/favoritemicro/dal"
	favorite "github.com/ClubWeGo/favoritemicro/kitex_gen/favorite/favoriteservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	dal.Init()

	registry, err := etcd.NewEtcdRegistry([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:1000")
	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "favoriteservice"}),
		server.WithRegistry(registry),
		server.WithServiceAddr(addr),
	)
	if err := svr.Run(); err != nil {
		log.Fatal(err)
	}
}
