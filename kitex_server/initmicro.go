package kitex_server

import (
	"log"

	"github.com/ClubWeGo/simple-tiktok/services/commentmicro/kitex_gen/comment/commentservice"
	"github.com/ClubWeGo/simple-tiktok/services/favoritemicro/kitex_gen/favorite/favoriteservice"
	"github.com/ClubWeGo/simple-tiktok/services/relationmicro/kitex_gen/relation/combineservice"
	"github.com/ClubWeGo/simple-tiktok/services/usermicro/kitex_gen/usermicro/userservice"
	"github.com/ClubWeGo/simple-tiktok/services/videomicro/kitex_gen/videomicro/videoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
)

var Userclient userservice.Client
var Videoclient videoservice.Client
var Relationclient combineservice.Client
var FavoriteClient favoriteservice.Client
var CommentClient commentservice.Client

func Init(r discovery.Resolver) {
	uc, err := userservice.NewClient("userservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	Userclient = uc

	vc, err := videoservice.NewClient("videoservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	Videoclient = vc

	rc, err := combineservice.NewClient("relationservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	Relationclient = rc

	fc, err := favoriteservice.NewClient("favoriteservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	FavoriteClient = fc

	Cc, err := commentservice.NewClient("commentservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	CommentClient = Cc
}
