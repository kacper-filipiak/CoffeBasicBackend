package model

import "github.com/go-pg/pg/v10"

type CoffeBag struct {
	ID        int64      `json:"id"`
	weight    int64      `json:"weight"`
	coffeType *CoffeType `json:"coffeType" pg:"rel:has-one"`
	roastDate int64      `json:"date"`
}

func GetAllCoffeBags(pgdb *pg.DB) ([]*CoffeBag, error) {
	coffeBags := make([]*CoffeBag, 0)

	err := pgdb.Model(&coffeBags).
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
