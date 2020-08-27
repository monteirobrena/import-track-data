package notification

import (
	"log"
)

func Error(currentError error) {
	if currentError != nil {
		log.Fatal("Error: ", currentError)
	}
}
