package rpc

import (
	"context"
	"github.com/ClubWeGo/usermicro/dal/model"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func GetUserByID(uid int64) (user model.User, err error) {
	r, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	client, err := userservice.NewClient("userservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.GetUserMethod(context.Background(), &usermicro.GetUserReq{
		Id: &uid,
	})
	//usermicro.CreateUserReq{Newuser_: &testuser, Password: "123456"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("创建一个新用户", resp)
	return model.User{}, nil
}
