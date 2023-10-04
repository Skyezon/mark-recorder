package datasources

import (
	"assignment-mezink/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// db singleton
var Db *sql.DB

func ConnectDB() (err error) {
	conf, err := utils.GetConfig()
	if err != nil {
		return utils.LogErr(err)
	}

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DB.DB_HOST, conf.DB.DB_PORT, conf.DB.DB_USER, conf.DB.DB_PASSWORD, conf.DB.DB_NAME)

	Db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return utils.LogErr(err, "fail to open database")
	}

	return nil
}
