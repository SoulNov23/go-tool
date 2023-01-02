package buffer

import (
	"context"
	"testing"
	"time"

	"github.com/SoulNov23/go-tool/pkg/unsafe"
)

func TestBuffer(t *testing.T) {
	lkBuffer := NewBuffer()

	write := make(chan struct{})
	read := make(chan struct{})

	timeout := 3 * time.Second
	ctx, cancel := context.WithCancel(context.Background())

	buf := unsafe.String2Byte("hello world")

	go func(ctx context.Context, lkBuffer *LinkedBuffer) {
		for {
			select {
			case <-ctx.Done():
				write <- struct{}{}
				return
			default:
				lkBuffer.Write(buf)
			}
		}
	}(ctx, lkBuffer)

	go func(ctx context.Context, lkBuffer *LinkedBuffer) {
		for {
			select {
			case <-ctx.Done():
				read <- struct{}{}
				return
			default:
				res, err := lkBuffer.Peek(40)
				if err != nil {
					t.Logf("LinkedBuffer.Peek: %v", err)
				} else {
					lkBuffer.Skip(cap(res))
					t.Logf("buf: %s", unsafe.Byte2String(res))
				}
				lkBuffer.GC()
			}
		}
	}(ctx, lkBuffer)

	go func(ctx context.Context, lkBuffer *LinkedBuffer) {
		for {
			select {
			case <-ctx.Done():
				read <- struct{}{}
				return
			default:
				res, err := lkBuffer.Read(40)
				if err != nil {
					t.Logf("LinkedBuffer.Read: %v", err)
				} else {
					t.Logf("buf: %s", unsafe.Byte2String(res))
				}
				lkBuffer.GC()
			}
		}
	}(ctx, lkBuffer)

	time.Sleep(timeout)

	DeleteBuffer(lkBuffer)

	cancel()

	<-write
	<-read
}
