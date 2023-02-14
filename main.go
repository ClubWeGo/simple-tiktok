package main

import (
	"github.com/ClubWeGo/usermicro/dal"
	usermicro "github.com/ClubWeGo/usermicro/kitex_gen/usermicro/userservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {

	dsn := "root:yutian@mysql+ssh(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
	dal.InitDB(dsn)

	r, err := etcd.NewEtcdRegistry([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:10000")
	svr := usermicro.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "userservice"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
