package util

import (
	"fmt"
	snow2 "github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/bwmarrin/snowflake"
)

var s, _ = snow2.NewSnowflake(int64(0), int64(0))
var node, _ = snowflake.NewNode(1)

func GetSnowflakeID() string {

	return node.Generate().String()
}

func GetSnowflakeInt() int64 {

	return node.Generate().Int64()
}

func GetSnowflakeID2() string {

	return fmt.Sprintf("%d", s.NextVal())
}
