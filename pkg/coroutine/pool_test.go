package coroutine

import (
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/soulnov23/go-tool/pkg/utils"
)

func TestStack(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			buffer := make([]byte, 10*1024)
			runtime.Stack(buffer, false)
			t.Logf("[PANIC] %v\n%s", err, utils.BytesToString(buffer))
			t.Logf("[PANIC] %v\n%s", err, utils.BytesToString(debug.Stack()))
		}
	}()
	panic("hello world")
}

func TestPool(t *testing.T) {
	pool, _ := NewPool(10, 10000, t.Errorf)
	for i := 0; i < 60; i++ {
		fn := func() {
			t.Logf("runtime.NumGoroutine: %d", runtime.NumGoroutine())
		}
		pool.Submit(fn)
	}
	pool.Release()
}
