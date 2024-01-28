package model

import (
	"github.com/go-pg/pg/v10"
    "fmt"
)

type CoffeType struct {
	tableName struct{} `pg:"coffe_types"`
    ID        int64    `json:"id" pg:",pk"`
	Name      string   `json:"name"`
}

func GetAllCoffeTypes(pgdb *pg.DB) ([]*CoffeType, error) {
	coffeTypes := make([]*CoffeType, 0)

	err := pgdb.Model(&coffeTypes).
		Select()

	return coffeTypes, err
}

func GetCoffeTypeByName(pgdb *pg.DB, name string) (*CoffeType, error) {
	coffeTypes := make([]*CoffeType, 0)

    fmt.Println(name)
	err := pgdb.Model(&coffeTypes).
        Where("coffe_type.name LIKE ?", name).
		Select()
    if err != nil {
      return nil, err
    }
	return coffeTypes[0], err
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
