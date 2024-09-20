package helper

import "github.com/google/uuid"

func Uuid() string {
	uuid := uuid.New()
	return uuid.String()
}

