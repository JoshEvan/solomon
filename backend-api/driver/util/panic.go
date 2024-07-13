package util

import (
	"log"
	"runtime/debug"
)

func PanicCapture() {
	if err := recover(); err != nil {
		log.Println("PANIC occurred:", err, string(debug.Stack()))
	}
}
