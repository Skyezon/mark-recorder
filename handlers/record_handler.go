package handlers

import (
	"assignment-mezink/services"
	"assignment-mezink/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// easier checking database content
func GetAllRecordHandler(w http.ResponseWriter, r *http.Request) {
	records, err := services.GetAllRecord()
	if err != nil {
		http.Error(w, utils.LogErr(err, "fail to get all records").Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(records)
	if err != nil {
		http.Error(w, utils.LogErr(err, "fail to marshall").Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetRecordHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, utils.LogErr(err, "Id must be number").Error(), http.StatusBadRequest)
		return
	}
	record, err := services.GetRecord(idInt)
	if err != nil {
		http.Error(w, utils.LogErr(err).Error(), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(record)
	if err != nil {
		http.Error(w, utils.LogErr(err).Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateRecordHandler(w http.ResponseWriter, r *http.Request) {
	var req utils.CreateRecordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, utils.LogErr(err, "Invalid request").Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = services.CreateRecord(utils.Record{
		Name:  req.Name,
		Marks: req.Marks,
		//CratedAt: time.Now().Format("2006-01-02 15:00:00"), manual input for easier testing
		CratedAt: req.CreatedAt,
	})
	if err != nil {
		http.Error(w, utils.LogErr(err, "Fail to create record").Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))

}

func GetSumRecordHandler(w http.ResponseWriter, r *http.Request) {
	var req utils.GetSumRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	records, err := services.GetSumRecords(req)
	if err != nil {
		http.Error(w, utils.LogErr(err, "Get sum record error").Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(utils.GetSumResponse{
		Code:    http.StatusOK,
		Msg:     "Success",
		Records: records,
	})
	if err != nil {
		http.Error(w, utils.LogErr(err, "Marshall error").Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
