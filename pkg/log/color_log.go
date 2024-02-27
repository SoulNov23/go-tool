package log

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	colorRed    = "\033[1;31m"
	colorGreen  = "\033[1;32m"
	colorYellow = "\033[1;33m"
	colorPurple = "\033[1;35m"
	colorWhite  = "\033[1;37m"
	colorReset  = "\033[m"
)

var (
	callers sync.Map
)

func ColorDebugf(formatter string, args ...any) {
	fmt.Printf("%s%s DEBUG %s %s%s\n", colorGreen, time.Now().Format(time.DateTime), caller(2), fmt.Sprintf(formatter, args...), colorReset)
}

func ColorInfof(formatter string, args ...any) {
	fmt.Printf("%s%s INFO %s %s%s\n", colorWhite, time.Now().Format(time.DateTime), caller(2), fmt.Sprintf(formatter, args...), colorReset)
}

func ColorWarnf(formatter string, args ...any) {
	fmt.Printf("%s%s WARN %s %s%s\n", colorYellow, time.Now().Format(time.DateTime), caller(2), fmt.Sprintf(formatter, args...), colorReset)
}

func ColorErrorf(formatter string, args ...any) {
	fmt.Printf("%s%s ERROR %s %s%s\n", colorRed, time.Now().Format(time.DateTime), caller(2), fmt.Sprintf(formatter, args...), colorReset)
}

func ColorFatalf(formatter string, args ...any) {
	fmt.Printf("%s%s FATAL %s %s%s\n", colorPurple, time.Now().Format(time.DateTime), caller(2), fmt.Sprintf(formatter, args...), colorReset)
}

// 返回调用者的"package/frame.File:frame.Line"
func caller(skip int) string {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+1, rpc[:])
	if n < 1 {
		return "unknown"
	}
	var frame runtime.Frame
	if f, ok := callers.Load(rpc[0]); ok {
		frame = f.(runtime.Frame)
	} else {
		frame, _ = runtime.CallersFrames(rpc).Next()
		callers.Store(rpc[0], frame)
	}
	strLine := strconv.Itoa(frame.Line)
	fullCaller := strings.Join([]string{frame.File, strLine}, ":")
	// 返回最后一个分隔符
	idx := strings.LastIndexByte(frame.File, '/')
	if idx == -1 {
		return fullCaller
	}
	// 返回倒数第二个分隔符
	idx = strings.LastIndexByte(frame.File[:idx], '/')
	if idx == -1 {
		return fullCaller
	}
	// 返回倒数第二个分隔符之后的所有内容
	return strings.Join([]string{frame.File[idx+1:], strLine}, ":")
}
