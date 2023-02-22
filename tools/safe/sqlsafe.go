package safe

import (
	"errors"
	"regexp"
)

// ref: https://blog.csdn.net/qq_40127376/article/details/108516561
var sqlInjectReg = regexp.MustCompile(`(.*\=.*\-\-.*)|(.*(\+|\-).*)|(.*\w+(%|\$|#|&)\w+.*)|(.*\|\|.*)|(.*\s+(and|or)\s+.*)|(.*\b(select|update|union|and|or|delete|insert|trancate|char|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b.*)`)

func SqlInjectCheck(input string) error {
	reg := sqlInjectReg.FindAllString(input, 1) // 匹配一个就行
	if reg != nil {
		return errors.New("输入存在非法字段")
	}
	return nil
}
