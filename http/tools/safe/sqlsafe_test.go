package safe

import (
	"log"
	"testing"
)

func TestSqlInjectCheck(t *testing.T) {
	str1 := "select 1"
	err := SqlInjectCheck(str1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("no")
}
