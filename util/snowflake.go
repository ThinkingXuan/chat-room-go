package util

import (
	snow2 "github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/bwmarrin/snowflake"
	"github.com/hashicorp/go-uuid"
	"time"
)

// 分布式ID,需要节点动态。
var s, _ = snow2.NewSnowflake(int64(time.Now().Nanosecond()%31), time.Now().UnixNano()%31)

var node, _ = snowflake.NewNode(1)

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

	return s.NextVal()
}
func GetSnowflakeID2() string {

	idx, _ := uuid.GenerateUUID()
	return idx
	//return fmt.Sprintf("%d", s.NextVal())
}
