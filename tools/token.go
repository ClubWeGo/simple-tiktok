package tools

import (
	"fmt"
)

var Tokens = "testtoken"
var UserID int64 = 2
var NextTime int64 = 1600000000

// 暂时简单生成
func GenerateToken(userid int64) string {
	token := Tokens + fmt.Sprint(userid)
	return token
}

// Token鉴权,暂时默认通过
func ValidateToken(token string) (valid bool, err error) {
	return true, nil // 这里交给后续token去做处理
}

func GetUserIdByToken(token string) (userid int64, err error) {
	// 1. 通过存储token的数据库查询
	// 2. 反解析token得到userid
	return UserID, nil
}

func GetNextTimeByToken(token string) (nextTime int64, err error) {
	// feed接口存储NextTime标识该用户的历史请求的时间，
	// 用户下滑刷新需要刷新这个时间戳
	// 可以通过时间比较来实现，超过几分钟之后，下次请求不使用该NextTime，使用最新时间
	return NextTime, nil
}
