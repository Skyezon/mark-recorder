package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
)

type Config struct {
	DB Database `json:"database"`
}

type Database struct {
	DB_NAME     string `json:"db_name"`
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
}

func GetConfig()(Config,error){
	content, err := ioutil.ReadFile("./env.json")
    if err != nil {
        return Config{},LogErr(err,"fail to read env.json")
    }

    var config Config
    err = json.Unmarshal(content,&config)

    if err != nil {
        return Config{},LogErr(err,"fail to unmarshall");
    }

    return config, nil

}

//easier error tracking & debuggin purposes
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
