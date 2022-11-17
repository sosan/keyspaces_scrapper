package utils

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

func GenerateRandomEmail() string {

	firstPart := rand.Int()
	secondPart := uuid.New()
	
	email := fmt.Sprintf("%d-%s@karenkey.com", firstPart, secondPart.String() ) 
	
	return email


}