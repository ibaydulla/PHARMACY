package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/utils"
)

type Medicinesfilter struct {
	Limit  int
	Offset int
	Search string
}

func LenStrMedicines(l []any) string {
	return strconv.Itoa(len(l))
}

func Medicineslist(c context.Context, f Medicinesfilter, moreArg ...int) ([]models.Medicines, error) {
	db := utils.GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlwhere := ``
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlwhere += ` and (name like '%$` + LenStrMedicines(sqlArgs) + `%') `

	}
	
	rows, err := db.Query(c, `select id, name , description, price, new_price, category_id
		from pharmacy_medicines
			where 1=1 `+sqlwhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}
	
	fmt.Println(`select id, name , description, price, new_price, category_id
		from pharmacy_medicines
			where 1=1 `+sqlwhere+`
		limit $1 offset $2`, sqlArgs)

	list := []models.Medicines{}
	for rows.Next() {
		item := models.Medicines{}
		rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.NewPrice, &item.CategoryID)
		list = append(list, item)
	}
	return list, nil
}

func MedicinesCreate(c context.Context, medicines models.Medicines) (models.Medicines, error) {

	_, err := utils.GetDB().Exec(c,
		"INSERT INTO pharmacy_medicines(id, name , description, price, new_price, category_id) VALUES ($1, $2, $3, $4, $5, $6)",
		medicines.ID, medicines.Name, medicines.Description, medicines.Price, medicines.NewPrice, medicines.CategoryID,
	)

	if err != nil {
		return models.Medicines{}, err
	}

	return medicines, nil
}

func MedicinesUpdate(c context.Context, id int, req models.Medicines) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`UPDATE users 
		 SET name=$1, description=$2, price=$3, new_price=$4, category_id=$5, 
		 WHERE id=$6`,
		req.ID, req.Name, req.Description, req.Price, req.NewPrice, req.CategoryID,
	)

	return err
}

func MedicinesDelete(c context.Context, id int) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`DELETE FROM pharmacy_medicines WHERE id=$1`,
		id,
	)

	return err
}
