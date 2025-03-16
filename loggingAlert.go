package logging

import (
	"fmt"
	"log"
)

func LogAttack(ip, data string) {
	logMessage := fmt.Sprintf("Suspicious request from %s: %s", ip, data)
	log.Println(logMessage)
}
