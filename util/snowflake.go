package util

import (
	"fmt"
	snow2 "github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/bwmarrin/snowflake"
)

var s, _ = snow2.NewSnowflake(int64(0), int64(0))

func GetSnowflakeID() string {

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return node.Generate().String()
}

func GetSnowflakeID2() string {

	return fmt.Sprintf("%d", s.NextVal())
}
