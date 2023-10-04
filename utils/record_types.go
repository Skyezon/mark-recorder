package utils

type Record struct {
	Id       int64   `json:"id"`
    Name     string  `json:"name"`
	Marks    []int64 `json:"marks"`
	CratedAt string  `json:"createdAt"`
}

type GetSumRequest struct {
	StartDate string
	EndDate   string
	MinCount  int
	MaxCount  int
}

type CreateRecordRequest struct {
	Name      string
	Marks     []int64
	CreatedAt string
}

//only for GetSumResponse
type ModRecord struct {
	Id         int64  `json:"id"`
	TotalMarks int64  `json:"totalMarks"`
	CreatedAt  string `json:"createdAt"`
}

type GetSumResponse struct {
    Code    int `json:"code"`
    Msg     string `json:"msg"`
    Records []ModRecord `json:"records"`
}
