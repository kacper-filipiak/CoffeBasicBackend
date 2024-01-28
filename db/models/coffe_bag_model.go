package model

import "github.com/go-pg/pg/v10"

type CoffeBag struct {
    tableName struct{} `pg:"coffe_bags"`
	ID        int64      `json:"id"`
	Weight    float64      `json:"weight"`
    CoffeTypeId int64   `pg:"coffe_type_id"` 
    CoffeType *CoffeType `json:"coffeType" pg:"rel:has-one,fk:coffe_type_id"`
	RoastDate int64      `json:"date"`
}

func GetAllCoffeBags(pgdb *pg.DB) ([]*CoffeBag, error) {
	coffeBags := make([]*CoffeBag, 0)

	err := pgdb.Model(&coffeBags).
        Relation("CoffeType").
		Select()

	return coffeBags, err
}

func CreateCoffeBag(db *pg.DB, req *CoffeBag) (*CoffeBag, error) {
	_, err := db.Model(req).Insert()
	if err != nil {
		return nil, err
	}

	coffeBag := &CoffeBag{}

	err = db.Model(coffeBag).
		Where("id = ?", req.ID).
		Select()

	return coffeBag, err
}
