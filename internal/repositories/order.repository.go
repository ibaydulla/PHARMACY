package repositories

import (
	"context"
	"fmt"

	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/utils"
)

type Orderfilter struct {
	Limit  int
	Offset int
	Search string
}

func Orderlist(c context.Context, f Orderfilter, moreArg ...int) ([]models.Order, error) {
	db := utils.GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlwhere := ``
	sqlArgs := []any{f.Limit, f.Offset}

	rows, err := db.Query(c, `select id, name , price, description
		from orders
			where 1=1 `+sqlwhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	fmt.Println(`select id, name , price, description
		from orders
			where 1=1 `+sqlwhere+`
		limit $1 offset $2`, sqlArgs)

	list := []models.Order{}
	for rows.Next() {
		item := models.Order{}
		rows.Scan(&item.ID, &item.Name, &item.Price, &item.Description)
		list = append(list, item)
	}
	return list, nil
}

func OrderCreate(c context.Context, order models.Order) (models.Order, error) {

	_, err := utils.GetDB().Exec(c,
		"INSERT INTO orders(id, name, price, description) VALUES ($1, $2, $3, $4)",
		order.ID, order.Name, order.Price, order.Description,
	)

	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func OrderUpdate(c context.Context, id int, req models.Order) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`UPDATE orders
		 SET name=$1, price=$2, description=$3, 
		 WHERE id=$4`,
		req.ID, req.Name, req.Price, req.Description,
	)

	return err
}

func OrderDelete(c context.Context, id int) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`DELETE FROM orders WHERE id=$1`,
		id,
	)

	return err
}
