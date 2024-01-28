package api

import (
	"encoding/json"
	model "kacperfilipiak/coffe_home/db/models"
	"log"

	"github.com/go-pg/pg/v10"

	"net/http"

)
type CreateCoffeTypeRequest struct {
	CoffeTypeName string `json:"Name"`
}

func createCoffeType(w http.ResponseWriter, r *http.Request) {
	req := &CreateCoffeTypeRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &CoffeTypeResponse{
			Success:    false,
			Error:      err.Error(),
			CoffeTypes: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error while sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &CoffeTypeResponse{
			Success:    false,
			Error:      "could not get the DB from context",
			CoffeTypes: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	coffeType, err := model.CreateCoffeType(pgdb, &model.CoffeType{Name: req.CoffeTypeName})
	if err != nil {
		res := &CoffeTypeResponse{
			Success:    false,
			Error:      err.Error(),
			CoffeTypes: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := &CoffeTypeResponse{
		Success:    true,
		Error:      "",
		CoffeTypes: []*model.CoffeType{coffeType},
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding after creating comment %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type CoffeTypeResponse struct {
	Success    bool               `json:"success"`
	Error      string             `json:"error"`
	CoffeTypes []*model.CoffeType `json:"coffeTypes"`
}

func getCoffeTypes(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &CoffeTypeResponse{
			Success:    false,
			Error:      "could not get DB from context",
			CoffeTypes: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	coffeTypes, err := model.GetAllCoffeTypes(pgdb)
	if err != nil {
		res := &CoffeTypeResponse{
			Success:    false,
			Error:      err.Error(),
			CoffeTypes: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &CoffeTypeResponse{
		Success:    true,
		Error:      "",
		CoffeTypes: coffeTypes,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comments: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
