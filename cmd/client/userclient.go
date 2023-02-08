package main

import (
	"context"
	"log"
	"time"

	"usermicro/kitex_gen/usermicro"
	"usermicro/kitex_gen/usermicro/userservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	client, err := userservice.NewClient("userservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	for {
		email := "test@qq.com"
		resp, err := client.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{Name: "hah", Email: &email, Password: "123456"})
		if err != nil {
			log.Fatal(err)
		}

		log.Println(resp)
		time.Sleep(time.Second * 3)
	}
}
