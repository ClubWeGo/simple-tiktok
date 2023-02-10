package main

import (
	"context"
	"log"

	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro/userservice"

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

		// // create user
		// email := "test@qq.com"
		// resp, err := client.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{Name: "hah", Email: &email, Password: "123456"})
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println(resp)

		// get user
		var id int64 = 1
		resp1, err := client.GetUserMethod(context.Background(), &usermicro.GetUserReq{Id: &id})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp1)
		// time.Sleep(time.Second * 3)

		// update user
		// var id int64 = 1
		// var newname = "hah"
		// var newemail = "1446@qq.com"
		// resp2, err := client.UpdateUserMethod(context.Background(), &usermicro.UpdateUserReq{Name: &newname, Email: &newemail})
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println(resp2)

		// // login user
		// // var id int64 = 1
		// var newname = "hah"
		// var newemail = "144611@qq.com"
		// resp3, err := client.UpdateUserMethod(context.Background(), &usermicro.UpdateUserReq{Name: &newname, Email: &newemail})
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// log.Println(resp3)

		break
	}
}
