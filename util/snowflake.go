package util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func GetSnowflakeID() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return node.Generate().String()
}
