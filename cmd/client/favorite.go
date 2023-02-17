package main

import (
	"context"
	"github.com/ClubWeGo/favoritemicro/kitex_gen/favorite"
	"github.com/ClubWeGo/favoritemicro/kitex_gen/favorite/favoriteservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func main() {
	resolver, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	//clt, err := favoriteservice.NewClient("favoriteservice", client.WithHostPorts("0.0.0.0:8888"))
	clt, err := favoriteservice.NewClient("favoriteservice", client.WithResolver(resolver))
	if err != nil {
		log.Fatal(err)
	}

	req := &favorite.FavoriteReq{UserId: 1, VideoId: 1, ActionType: 1}
	res, err := clt.FavoriteMethod(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
