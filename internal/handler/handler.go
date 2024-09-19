package handler

import (
	"fmt"
)

func StartProcessing(msgCh <-chan []byte) {
	for msg := range msgCh {
		// Обработка сообщения
		fmt.Printf("Received message: %s\n", msg)
	}
}
