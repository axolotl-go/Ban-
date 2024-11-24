package services

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateCardNumber() string {
	rand.Seed(time.Now().UnixNano())
	cardNumber := "9510"
	for i := 0; i < 12; i++ {
		cardNumber += fmt.Sprintf("%d", rand.Intn(10))
	}

	return cardNumber
}

func GeneratePin() string {
	rand.Seed(time.Now().UnixNano())
	cardPin := ""

	for i := 0; i < 4; i++ {
		cardPin += fmt.Sprintf("%d", rand.Intn(10))
	}

	return cardPin
}
