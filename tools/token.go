package tools

import "fmt"

var Tokens = "testtoken"

// 暂时简单生成
func GetToken(userid int64) string {
	token := Tokens + fmt.Sprint(userid)
	return token
}
