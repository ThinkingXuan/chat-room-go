package pool

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/model/mysql"
	"sync"
)

var (
	goPool *Pool
	wg     = sync.WaitGroup{}
)

func InitGoRoutinePool(maxConn uint64) bool {
	goPool, _ = NewPool(maxConn)
	return true
}

// Work 使用协程池
func Work(roomID, roomName string) {
	wg.Add(1)
	goPool.Put(&Task{
		Handler: func(v ...interface{}) {
			mysql.CreateAsyncRoom(roomID, roomName)
			wg.Done()

		},
		Params: []interface{}{roomID, roomName},
	})
	wg.Wait()
}

func WorkSendMessage(req *rr.ReqMessage) {
	wg.Add(1)
	goPool.Put(&Task{
		Handler: func(v ...interface{}) {
			mysql.CreateSyncMessage(req)
			wg.Done()

		},
		Params: []interface{}{},
	})
	wg.Wait()
}
