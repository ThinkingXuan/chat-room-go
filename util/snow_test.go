package util

import (
	"fmt"
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
