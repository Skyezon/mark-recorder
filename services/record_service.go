package services

import (
	"assignment-mezink/datasources"
	"assignment-mezink/utils"
)

func GetRecord(id int)(utils.Record, error) {
    record ,err := datasources.GetRecord(id)
    if err != nil {
        return utils.Record{},utils.LogErr(err)
    }
    return record,nil
}

func CreateRecord(newRecord utils.Record)error {
    err := datasources.InsertRecord(newRecord)
    if err != nil {
        return utils.LogErr(err)
    }
    return nil
}

func GetSumRecords(req utils.GetSumRequest) ([]utils.Record, error) {
	records, err := datasources.GetRecordBetweenDates(req.StartDate, req.EndDate)
	if err != nil {
		return []utils.Record{}, utils.LogErr(err)
	}
	result := make([]utils.Record, 0)
	for _, record := range records {
		if checkSumBetween(record, req.MinCount, req.MaxCount) {
			result = append(result, record)
		}
	}
	return result, nil
}

func checkSumBetween(record utils.Record, min int, max int) bool {
	sum := 0
	for num := range record.Marks {
		sum += num
		if sum > max {
			return false
		}
	}
	if sum < min {
		return false
	}
	return true
}
