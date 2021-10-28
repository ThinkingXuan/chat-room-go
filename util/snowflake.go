package util

import (
	snow2 "github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/bwmarrin/snowflake"
	"time"
)

// 分布式ID,需要节点动态产生。
var s, _ = snow2.NewSnowflake(int64(time.Now().Nanosecond()%31), time.Now().UnixNano()%31)

// 分布式ID,需要节点动态产生。
var node, _ = snowflake.NewNode(GetLocalIntShortIP())

//func GetSnowflakeID() string {
//
//	return node.Generate().String()
//}
//
//func GetSnowflakeInt() int64 {
//
//	return node.Generate().Int64()
//}

func GetSnowflakeInt2() int64 {
	return node.Generate().Int64()
}
func GetSnowflakeID2() string {
	return node.Generate().String()
}
