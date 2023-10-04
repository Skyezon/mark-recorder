package datasources

import (
	"assignment-mezink/utils"
	"fmt"
)

func InsertRecord(newRecord utils.Record) error {
	query := fmt.Sprintf("INSERT INTO records VALUES ($1,$2,$3);")
	stmt, err := Db.Prepare(query)
	if err != nil {
		return utils.LogErr(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(newRecord.Name, newRecord.Marks, newRecord.CratedAt)
	if err != nil {
		return err
	}
	return nil
}

func GetRecord(id int) (utils.Record, error) {
	query := fmt.Sprintf("SELECT * FROM records WHERE id = $1;")
	stmt, err := Db.Prepare(query)
	if err != nil {
		return utils.Record{}, utils.LogErr(err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	var record utils.Record
	if err = row.Scan(&record.Id, &record.Name, &record.Marks, &record.CratedAt); err != nil {
		return utils.Record{}, utils.LogErr(err)
	}
	return record, nil
}

func GetRecordBetweenDates(startDate string, endDate string) ([]utils.Record, error) {
	query := fmt.Sprintf("SELECT * FROM records WHERE createdAt BETWEEN $1 AND $2;")
	stmt, err := Db.Prepare(query)
	if err != nil {
		return []utils.Record{}, utils.LogErr(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(startDate, endDate)
	if err != nil {
		return []utils.Record{}, utils.LogErr(err)
	}

	result := make([]utils.Record, 0)

	for rows.Next() {
		var record utils.Record
		if err := rows.Scan(&record.Id, &record.Name, &record.Marks, &record.CratedAt); err != nil {
			return []utils.Record{}, utils.LogErr(err)
		}
		result = append(result, record)
	}

	return result, nil
}
