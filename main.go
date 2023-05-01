package logger

import (
	"log"
)

func Loginfo(message string) {
	log.Printf("Info %v", message)
}
