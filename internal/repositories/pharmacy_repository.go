package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/utils"
)

type Pharmacyfilter struct {
	Limit  int
	Offset int
	Search string
}

func Len(l []any) string {
	return strconv.Itoa(len(l))
}

func Pharmacylist(c context.Context, f Pharmacyfilter, moreArg ...int) ([]models.Pharmacy, error) {
	db := utils.GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlwhere := ``
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlwhere += ` and (name like '%$` + Len(sqlArgs) + `%') `

	}

	rows, err := db.Query(c, `select id, name , address, pharmacy_hours
		from pharmacy_db
			where 1=1 `+sqlwhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	fmt.Println(`select id, name , address, pharmacy_hours
		from pharmacy_db
			where 1=1 `+sqlwhere+`
		limit $1 offset $2`, sqlArgs)

	list := []models.Pharmacy{}
	for rows.Next() {
		item := models.Pharmacy{}
		rows.Scan(&item.ID, &item.Name, &item.Address, &item.Pharmacyhours)
		list = append(list, item)
	}
	return list, nil
}

func PharmacyCreate(c context.Context, pharmacy models.Pharmacy) (models.Pharmacy, error) {

	_, err := utils.GetDB().Exec(c,
		"INSERT INTO users(id, name, address, pharmacy_hours) VALUES ($1, $2, $3, $4)",
		pharmacy.ID, pharmacy.Name, pharmacy.Address, pharmacy.Pharmacyhours,
	)

	if err != nil {
		return models.Pharmacy{}, err
	}

	return pharmacy, nil
}

func PharmacyUpdate(c context.Context, id int, req models.Pharmacy) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`UPDATE pharmacies 
		 SET name=$1, address=$2, pharmacy_hours=$3
		 WHERE id=$4`,
		req.ID, req.Name, req.Address, req.Pharmacyhours,
	)

	return err
}

func PharmacyDelete(c context.Context, id int) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`DELETE FROM pharmacies WHERE id=$1`,
		id,
	)

	return err
}
