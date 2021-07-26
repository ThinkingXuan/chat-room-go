package pool

import (
	"testing"
)

func TestPool(t *testing.T) {
	InitGoRoutinePool(1000)
	for i := 0; i < 1000; i++ {
		Work("123", "123")
	}
}
