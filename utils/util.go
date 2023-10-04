package utils

import (
	"fmt"
	"log"
	"runtime"
)

type Config struct {
	DB Database `json:"database"`
}

type Database struct {
    DB_NAME     string `json:"db_name"`
    DB_PASSWORD string `json:"db_password"`
    DB_HOST     string `json:"db_host"`
    DB_PORT     string `json:"db_port"`
    DB_USER     string `json:"db_user"`
}


//easier error tracking & debuggin purposes: basicly will show line of error where it happen & optional error message (from internet)
func LogErr(err error, msgToDev ...string) error {
	pc := make([]uintptr, 5)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	log.Println(err)
	if len(msgToDev) > 0 {
		log.Printf("debug msg : %s", msgToDev)
	}
	for {
		frame, more := frames.Next()
		log.Printf("%s:%d", frame.File, frame.Line)
		if !more {
			break
		}
	}
	fmt.Println("")
	return err
}
