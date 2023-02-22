package rpc

import (
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/kitex_gen/favorite/favoriteservice"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro/videoservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var VideoClient videoservice.Client
var FavoriteClient favoriteservice.Client

func InitRPC() {
	resolver, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	VideoClient, err = videoservice.NewClient("videoservice", client.WithResolver(resolver))
	if err != nil {
		log.Fatal(err)
	}

	FavoriteClient, err = favoriteservice.NewClient("favoriteservice", client.WithResolver(resolver))
	if err != nil {
		log.Fatal(err)
	}

}
