package benchmark

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/config"
	"chat-room-go/model/mysql"
	"chat-room-go/util"
	"fmt"
	"github.com/golang/glog"
	"testing"
)

func initMy(x, y int) {
	var cfgPath = "D:\\workspace\\go_workspace\\src\\chat-room-go\\config\\conf\\config.yaml"

	err := config.InitConfig(cfgPath)
	if err != nil {
		glog.Error(err)
		panic(err)
	}
	//初始化
	if err := mysql.InitMySql(); err != nil {
		glog.Error(err)
		panic("数据库初始化失败")
	}
	//defer model.Close()

	//模型绑定
	mysql.InitDBTable(x, y)
}

func Write() {
	req := &rr.ReqUser{
		Username:  util.GetSnowflakeID2(),
		FirstName: "234234",
		LastName:  "234324",
		Email:     "2234234",
		Password:  "234324234",
		Phone:     "234324",
	}
	err := mysql.CreateUser(req)
	if err != nil {
		fmt.Println(err)
	}
}

func BenchmarkWrite1_1(b *testing.B) {
	initMy(1, 1)
	for n := 0; n < b.N; n++ {
		Write()
	}
}
func BenchmarkWrite2_1(b *testing.B) {
	initMy(2, 1)
	for n := 0; n < b.N; n++ {
		Write()
	}
}
func BenchmarkWrite3_1(b *testing.B) {
	initMy(3, 1)
	for n := 0; n < b.N; n++ {
		Write()
	}
}

func BenchmarkWrite4_2(b *testing.B) {
	initMy(4, 2)
	for n := 0; n < b.N; n++ {
		Write()
	}
}
func BenchmarkWrite6_3(b *testing.B) {
	initMy(6, 3)
	for n := 0; n < b.N; n++ {
		Write()
	}
}
func BenchmarkWrite8_4(b *testing.B) {
	initMy(8, 4)
	for n := 0; n < b.N; n++ {
		Write()
	}
}

func BenchmarkWrite12_6(b *testing.B) {
	initMy(12, 6)
	for n := 0; n < b.N; n++ {
		Write()
	}
}
func BenchmarkWrite14_7(b *testing.B) {
	initMy(14, 7)
	for n := 0; n < b.N; n++ {
		Write()
	}
}

func BenchmarkWrite18_9(b *testing.B) {
	initMy(18, 9)
	for n := 0; n < b.N; n++ {
		Write()

	}
}

func BenchmarkWrite20_10(b *testing.B) {
	initMy(20, 10)
	for n := 0; n < b.N; n++ {
		Write()
	}
}

func BenchmarkWrite30_15(b *testing.B) {
	initMy(30, 15)
	for n := 0; n < b.N; n++ {
		Write()
	}
}

func BenchmarkWrite32_16(b *testing.B) {
	initMy(32, 16)
	for n := 0; n < b.N; n++ {
		Write()
	}
}

func BenchmarkWrite40_20(b *testing.B) {
	initMy(40, 20)
	for n := 0; n < b.N; n++ {
		Write()
	}
}
