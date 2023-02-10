package tools

import "fmt"

var Tokens = "testtoken"
var UserID int64 = 2

// 暂时简单生成
func GetToken(userid int64) string {
	token := Tokens + fmt.Sprint(userid)
	return token
}

// Token鉴权,暂时默认通过
func ValidateToken(token string) (userid int64, ifValidToken bool) {
	return UserID, true // 这里交给后续token去做处理
}
