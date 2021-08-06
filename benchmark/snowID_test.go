package benchmark

import (
	"chat-room-go/util"
	"testing"
)

func BenchmarkSnowID(b *testing.B) {

	for n := 0; n < b.N; n++ {
		 util.GetSnowflakeID()
	}
}

func BenchmarkSnowID2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		util.GetSnowflakeID2()
	}
}
