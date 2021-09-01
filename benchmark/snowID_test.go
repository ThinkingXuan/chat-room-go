package benchmark

import (
	"chat-room-go/util"
	"fmt"
	snow2 "github.com/GUAIK-ORG/go-snowflake/snowflake"
	"strings"
	"testing"
	"time"
)

func BenchmarkSnowID(b *testing.B) {

}

func BenchmarkSnowID2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		util.GetSnowflakeID2()
	}
}

func TestSnowID(t *testing.T) {
	//  datacenterId workerId范围为[0,31]
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		fmt.Println(int64(time.Now().Nanosecond() % 31))
		fmt.Println(time.Now().UnixNano() % 31)
	}
	_, err := snow2.NewSnowflake(int64(time.Now().Nanosecond()%31), int64(time.Now().Nanosecond()%31))
	fmt.Println(err)
}

func TestName(t *testing.T) {
	ip := "192.16.1.104:"
	ip = strings.Split(ip, ":")[0]
	fmt.Println(ip)
}
