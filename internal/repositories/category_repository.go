package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ibaydulla/internal/models"
	"github.com/ibaydulla/internal/utils"
)

type Categoryfilter struct {
	Limit  int
	Offset int
	Search string
}

func LenStrcategory(l []any) string {
	return strconv.Itoa(len(l))
}

func Categorylist(c context.Context, f Categoryfilter, moreArg ...int) ([]models.Category, error) {
	db := utils.GetDB()
	if f.Limit == 0 {
		f.Limit = 10
	}
	sqlwhere := ``
	sqlArgs := []any{f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlwhere += ` and (name like '%$` + LenStrcategory(sqlArgs) + `%') `

	}

	rows, err := db.Query(c, `select id, name 
		from categories
			where 1=1 `+sqlwhere+`
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	fmt.Println(`select id, name 
		from categories
			where 1=1 `+sqlwhere+`
		limit $1 offset $2`, sqlArgs)

	list := []models.Category{}
	for rows.Next() {
		item := models.Category{}
		rows.Scan(&item.ID, &item.Name)
		list = append(list, item)
	}
	return list, nil
}

func CategoryCreate(c context.Context, category models.Category) (models.Category, error) {

	_, err := utils.GetDB().Exec(c,
		"INSERT INTO categories(id, name) VALUES ($1, $2)",
		category.ID, category.Name,
	)

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func CategoryUpdate(c context.Context, id int, req models.Category) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`UPDATE categories 
		 SET name=$1 
		 WHERE id=$2`,
		req.ID, req.Name,
	)

	return err
}

func CategoryDelete(c context.Context, id int) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`DELETE FROM categories WHERE id=$1`,
		id,
	)

	return err
}
