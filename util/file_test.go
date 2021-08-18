package util

import (
	"fmt"
	"strings"
	"testing"
)

func TestFileWrite(t *testing.T) {
	host := "\n192.168.1.104:123124\n192.168.1.102:123124\n192.168.1.103:123124"
	WriteWithFile("./ip-address", host)
}

func TestReadWithFile(t *testing.T) {
	host := ReadWithFile("./ip-address")
	h := strings.Split(host, "\n")
	fmt.Println(h, len(h))
}
