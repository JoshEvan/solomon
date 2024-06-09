package util

import (
	"log"
)

func PanicCapture() {
	if err := recover(); err != nil {
		log.Println("PANIC occurred:", err)
	}
}
