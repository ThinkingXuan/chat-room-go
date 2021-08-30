package util

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"log"
	"testing"
)

func TestGetSnowflakeID(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GetSnowflakeInt())
	}
}

func TestGetSnowflakeID2(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Println(GetSnowflakeID2())

	}
}

func TestReg(t *testing.T) {
	//reg := regexp.MustCompile("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
	//emailOk := reg.MatchString("jing54@yahoo.com")
	//log.Println(emailOk)
	//regEmail     := regexp.MustCompile("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
	//regPhone     := regexp.MustCompile("^1(3[0-9]|5[0-3,5-9]|7[1-3,5-8]|8[0-9])\\d{8}$")
	//regUserName  := regexp.MustCompile("^[\\\\u4e00-\\\\u9fa5_a-zA-Z0-9-]{1,16}$")
	regPassword := regexp2.MustCompile("^(?![a-zA-Z]+$)(?!\\d+$)(?![!@#$%^&*]+$)[a-zA-Z\\d!@#()_$%-^.&*]+$", regexp2.RE2)
	//regFirstName := regexp.MustCompile("^[\\u4e00-\\u9fa5_a-zA-Z]+$")
	//regLastName  := regexp.MustCompile("^[\\u4e00-\\u9fa5_a-zA-Z]+$")

	usrenameOk, _ := regPassword.MatchString("-_4o^TBF!Sc%w123")
	log.Println(usrenameOk)
}
