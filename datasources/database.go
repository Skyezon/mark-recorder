package datasources

import (
	"assignment-mezink/utils"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// db singleton
var Db *sql.DB

func ConnectDB() (err error) {
	psqlconn := os.Getenv("CONNECTION_STRING") //from docker
    if psqlconn == ""{
        log.Fatal("CONNECTION_STRING not set")
    }

	Db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return utils.LogErr(err, "fail to open database")
	}

	return nil
}
