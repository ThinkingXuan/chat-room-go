package util

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func GetSnowflakeID() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return node.Generate().Int64()
}
