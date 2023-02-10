package main

import (
	"context"
	"log"
	"time"

	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro/userservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type User struct {
	Name     string
	Password string
	Email    string
}

func generateTestData() []User {
	data := make([]User, 0)
	// Thanks https://www.jianshu.com/p/d5f00ad58572

	data = append(data, User{
		"testuser1", "123456", "seclee@cc.com",
	})
	data = append(data, User{
		"testuser2", "123456", "1446103183@qq.com",
	})
	data = append(data, User{
		"testuser3", "123456", "seclee@126.com",
	})
	return data
}

func main() {
	r, err := etcd.NewEtcdResolver([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	client, err := userservice.NewClient("userservice", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}

	// create user
	datalist := generateTestData()

	for _, user := range datalist {
		resp, err := client.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{Name: user.Name, Email: &user.Email, Password: user.Password})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second * 2)
	}

	// // get user
	// var id int64 = 1
	// resp1, err := client.GetUserMethod(context.Background(), &usermicro.GetUserReq{Id: &id})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp1)
	// // time.Sleep(time.Second * 3)

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
}
