package utils

import "log"

func CheckErrors(err error, code, msg string) {
	if err != nil {
		log.Fatalf("Code: %v Error %v", code, msg)
	}
}
