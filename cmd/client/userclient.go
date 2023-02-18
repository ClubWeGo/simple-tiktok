package main

import (
	"context"
	"log"

	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro"
	"github.com/ClubWeGo/usermicro/kitex_gen/usermicro/userservice"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type User struct {
	Name      string
	Password  string
	Email     string
	Signature string
}

func generateTestData() []User {
	data := make([]User, 0)
	// Thanks https://www.jianshu.com/p/d5f00ad58572

	data = append(data, User{
		"testuser1", "123456", "seclee@cc.com", "我是testuser1",
	})
	data = append(data, User{
		"testuser2", "123456", "1446103183@qq.com", "美好生活",
	})
	data = append(data, User{
		"testuser3", "123456", "seclee@126.com", "我是testuser3",
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

	// // create user
	// datalist := generateTestData()

	// for _, user := range datalist {
	// 	newuser := usermicro.CreateUserInfo{
	// 		Name:      user.Name,
	// 		Email:     &user.Email,
	// 		Signature: &user.Signature,
	// 	}
	// 	resp, err := client.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{Newuser_: &newuser, Password: user.Password})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(resp)
	// 	time.Sleep(time.Second * 2)
	// }

	var email, signature string
	testuser := usermicro.CreateUserInfo{
		Name:      "abc",
		Email:     &email,
		Signature: &signature,
	}
	resp, err := client.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{Newuser_: &testuser, Password: "123456"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("创建一个新用户", resp)

	resp1, err := client.CreateUserMethod(context.Background(), &usermicro.CreateUserReq{Newuser_: &testuser, Password: "123456"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("重复创建一个新用户", resp1)

	// get user
	var id1 int64 = 1
	resp2, err := client.GetUserMethod(context.Background(), &usermicro.GetUserReq{Id: &id1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("查询用户", resp2)

	// update user
	var id2 int64 = 1
	var newpassword = "654321"
	var updateuser = usermicro.UpdateUserInfo{
		Password: &newpassword,
	}
	resp3, err := client.UpdateUserMethod(context.Background(), &usermicro.UpdateUserReq{Id: id2, UpdateData: &updateuser})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("更新用户信息", resp3)

	// login user
	// var id int64 = 1
	var newname = "testuser1"
	var Password = "654321"
	resp4, err := client.LoginUserMethod(context.Background(), &usermicro.LoginUserReq{Name: newname, Password: Password})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("正确登录信息", resp4)

	var newname1 = "testuser1"
	var Password1 = "123456"
	resp5, err := client.LoginUserMethod(context.Background(), &usermicro.LoginUserReq{Name: newname1, Password: Password1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("错误登录信息", resp5)

	// get user_set by id_set
	resp6, err := client.GetUserSetByIdSetMethod(context.Background(), &usermicro.GetUserSetByIdSetReq{
		IdSet: []int64{1, 2, 3},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("批量查询用户信息", resp6)

	// update relation cache
	resp7, err := client.UpdateRelationMethod(context.Background(), &usermicro.UpdateRelationCacheReq{
		Id: 1,
		NewData_: &usermicro.UpdateRelationCache{
			FollowCount:   1,
			FollowerCount: 2,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("更新用户信息缓存", resp7)
}
