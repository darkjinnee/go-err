package goerr

import "log"

func Fatal(err error, msg string) {
	if err != nil {
		log.Fatalf("%s...\n%s", msg, err)
	}
}

func Log(err error, msg string) {
	if err != nil {
		log.Printf("%s...\n%s", msg, err)
	}
}
