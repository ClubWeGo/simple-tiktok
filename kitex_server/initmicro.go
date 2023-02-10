package kitex_server

import (
	"log"

	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro/userservice"
	"github.com/ClubWeGo/videomicro/kitex_gen/videomicro/videoservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
)

var Userclient userservice.Client
var Videoclient videoservice.Client

func InitMicroServer(r discovery.Resolver) {
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
}
