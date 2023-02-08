package main

import (
	"log"
	"usermicro/kitex_gen/usermicro/userservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	r, err := etcd.NewEtcdRegistry([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userservice"}),
		server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
