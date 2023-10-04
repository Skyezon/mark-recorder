package utils

type Record struct {
	Id       int64
	Name     string
	Marks    []int
	CratedAt string
}

type GetSumRequest struct {
	StartDate string
	EndDate   string
	MinCount  int
	MaxCount  int
}

type CreateRecordRequest struct {
    Name string
    Marks []int
}

type GetSumResponse struct {
    Code int
    Msg string
    Records []Record
}
