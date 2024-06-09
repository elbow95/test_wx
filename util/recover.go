package util

import (
	"log"
	"runtime"
)

func DefaultRecover() {
	if e := recover(); e != nil {
		const size = 64 << 10
		buf := make([]byte, size)
		buf = buf[:runtime.Stack(buf, false)]
		log.Fatalf("goroutine panic: %s: %s", e, buf)
	}
}

func GoWithRecovery(df func(), f func()) {
	go func() {
		defer df()
		f()
	}()
}

func GoWithDefaultRecovery(f func()) {
	GoWithRecovery(DefaultRecover, f)
}
