package repositories

import (
	"context"
	"strconv"

)

type Userfilter struct{
	Limit int 
	Offset int 
	Search string 
	Role string
}

func LenStr(l []any) string {
	return strconv.Itoa(len(l))
}

func Userlist(c context.Context, f Userfilter, moreArg ...int) ([]models.User, error) {
	db := utils.GetDB()
	sqlwhere := ` `
	sqlArgs := []any{f,f.Limit, f.Offset}
	if f.Search != "" {
		sqlArgs = append(sqlArgs, f.Search)
		sqlwhere += ` and (first_name ilike '%$` + LenStr(sqlArgs) + `%') `

	}
	if f.Role != "" {
		sqlArgs = append(sqlArgs, f.Role)
		sqlwhere += ` and (role=$` + LenStr(sqlArgs) + `)`

	}
	 
	rows, err := db.Query(c, `select id, first_name, last_name, role, password, email
		from users
			where 1=1 ` +sqlWhere+ `
		limit $1 offset $2`, sqlArgs...)
	if err != nil {
		return nil, err
	}

	list := []models.User{}
	for rows.Next() {
		item := models.User{}
		rows.Scan(&item.ID, &item.Firstname, &item.Lastname, &item.Role, &item.Password, &item.Email)
		list = append(list, item)
	}
	return list nil
}

func UserCreate(c context.Context, user models.User) (models.User, error) {

	_, err := utils.GetDB().Exec(context.Background(),
		"INSERT INTO users(id, name, email, password, role) VALUES ($1,$2,$3,$4,$5)",
		user.ID, user.Name, user.Email, user.Password, user.Role,
	)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UserUpdate(c context.Context, id int, req models.User) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`UPDATE users 
		 SET name=$1, email=$2, password=$3, role=$4 
		 WHERE id=$5`,
		req.Name, req.Email, req.Password, req.Role, id,
	)

	return err
}


func UserDelete(c context.Context, id int) error {
	db := utils.GetDB()

	_, err := db.Exec(c,
		`DELETE FROM users WHERE id=$1`,
		id,
	)

	return err
}

