package model

import (
	"github.com/go-pg/pg/v10"
)

type CoffeType struct {
	tableName struct{} `pg:"coffe_types"`
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
}

func GetAllCoffeTypes(pgdb *pg.DB) ([]*CoffeType, error) {
	coffeTypes := make([]*CoffeType, 0)

	err := pgdb.Model(&coffeTypes).
		Select()

	return coffeTypes, err
}

func CreateCoffeType(db *pg.DB, req *CoffeType) (*CoffeType, error) {
	_, err := db.Model(req).Insert()
	if err != nil {
		return nil, err
	}

	coffeType := &CoffeType{}

	err = db.Model(coffeType).
		Where("coffe_type.id = ?", req.ID).
		Select()

	return coffeType, err
}
