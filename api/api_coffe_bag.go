package api

import (
	"encoding/json"
	model "kacperfilipiak/coffe_home/db/models"
	"log"

    "fmt"
	"github.com/go-pg/pg/v10"

	"net/http"
    "time"
)

type CoffeBagResponse struct {
	Success    bool               `json:"success"`
	Error      string             `json:"error"`
	CoffeBags []*model.CoffeBag `json:"coffeBags"`
}

type CreateCoffeBagRequest struct {
	CoffeTypeName string `json:"TypeName"`
    Weight float64  `jsno:"Weight"`
    RoastDate string `json:"RoastDate"`
}

func createCoffeBag(w http.ResponseWriter, r *http.Request) {
	req := &CreateCoffeBagRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		res := &CoffeBagResponse{
			Success:    false,
			Error:      err.Error(),
			CoffeBags: nil,
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
		res := &CoffeBagResponse{
			Success:    false,
			Error:      "could not get the DB from context",
			CoffeBags: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
    roastDate, err := time.Parse("02/01/2006", req.RoastDate)
    coffeType, err := model.GetCoffeTypeByName(pgdb, req.CoffeTypeName) 
    bag := model.CoffeBag{Weight: req.Weight, RoastDate: roastDate.Unix(), CoffeTypeId: coffeType.ID}
    coffeBag, err := model.CreateCoffeBag(pgdb, &bag)
    fmt.Println(coffeBag)
	if err != nil {
		res := &CoffeBagResponse{
			Success:    false,
			Error:      err.Error(),
			CoffeBags: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := &CoffeBagResponse{
		Success:    true,
		Error:      "",
		CoffeBags: []*model.CoffeBag{coffeBag},
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding after creating comment %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}


func getCoffeBags(w http.ResponseWriter, r *http.Request) {
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &CoffeBagResponse{
			Success:    false,
			Error:      "could not get DB from context",
			CoffeBags: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	coffeBags, err := model.GetAllCoffeBags(pgdb)
	if err != nil {
		res := &CoffeBagResponse{
			Success:    false,
			Error:      err.Error(),
			CoffeBags: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("error sending response %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := &CoffeBagResponse{
		Success:    true,
		Error:      "",
		CoffeBags: coffeBags,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("error encoding comments: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
