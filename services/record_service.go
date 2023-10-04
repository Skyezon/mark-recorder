package services

import (
	"assignment-mezink/datasources"
	"assignment-mezink/utils"
)

func GetAllRecord() ([]utils.Record, error) {
	records, err := datasources.GetAllRecord()
	if err != nil {
		return []utils.Record{}, utils.LogErr(err)
	}
	return records, nil
}

func GetRecord(id int) (utils.Record, error) {
	record, err := datasources.GetRecord(id)
	if err != nil {
		return utils.Record{}, utils.LogErr(err)
	}
	return record, nil
}

func CreateRecord(newRecord utils.Record) error {
	err := datasources.InsertRecord(newRecord)
	if err != nil {
		return utils.LogErr(err)
	}
	return nil
}

func GetSumRecords(req utils.GetSumRequest) ([]utils.ModRecord, error) {
	records, err := datasources.GetRecordBetweenDates(req.StartDate, req.EndDate)
	if err != nil {
		return []utils.ModRecord{}, utils.LogErr(err)
	}
	result := make([]utils.ModRecord, 0)
	for _, record := range records {
		isSumAsCriteria, sum := checkSumBetween(record, int64(req.MinCount), int64(req.MaxCount))
		if isSumAsCriteria {
			result = append(result, utils.ModRecord{
				Id:         record.Id,
				TotalMarks: sum,
				CreatedAt:  record.CratedAt,
			})
		}
	}
	return result, nil
}

func checkSumBetween(record utils.Record, min int64, max int64) (bool, int64) {
	sum := int64(0)
	for _, num := range record.Marks {
		sum += num
		if sum > max {
			return false, 0
		}
	}
	if sum < min {
		return false, 0
	}
	return true, sum
}
