package repositories

import (
"context"
"fmt"
"strconv"

"github.com/ibaydulla/internal/models"
"github.com/ibaydulla/internal/utils"

)

type Userfilter struct {
Limit  int
Offset int
Search string
Role   string
}

func LenStr(l []any) string {
return strconv.Itoa(len(l))
}

func Userlist(c context.Context, f Userfilter, moreArg ...int) ([]models.User, error) {

db := utils.GetDB()

if f.Limit == 0 {
	f.Limit = 10
}

sqlwhere := ``
sqlArgs := []any{f.Limit, f.Offset}

if f.Role != "" {
	sqlArgs = append(sqlArgs, f.Role)
	sqlwhere += ` and role=$` + LenStr(sqlArgs)
}

rows, err := db.Query(
	c,
	`SELECT id, name, role, password, email
	 FROM users
	 WHERE 1=1 `+sqlwhere+`
	 LIMIT $1 OFFSET $2`,
	sqlArgs...,
)

if err != nil {
	return nil, err
}

defer rows.Close()

fmt.Println(sqlArgs)

list := []models.User{}

for rows.Next() {

	var item models.User

	err := rows.Scan(
		&item.ID,
		&item.Name,
		&item.Role,
		&item.Password,
		&item.Email,
	)

	if err != nil {
		return nil, err
	}

	list = append(list, item)
}

return list, nil

}

func GetUserByEmail(c context.Context, email string) (models.User, error) {

var user models.User

err := utils.GetDB().QueryRow(
	c,
	`SELECT id, name, role, password, email
	 FROM users
	 WHERE email=$1`,
	email,
).Scan(
	&user.ID,
	&user.Name,
	&user.Role,
	&user.Password,
	&user.Email,
)

if err != nil {
	return models.User{}, err
}

return user, nil

}

func UserCreate(c context.Context, user models.User) (models.User, error) {

_, err := utils.GetDB().Exec(
	c,
	`INSERT INTO users
	(name, email, password, role)
	VALUES ($1, $2, $3, $4)`,
	user.Name,
	user.Email,
	user.Password,
	user.Role,
)

if err != nil {
	return models.User{}, err
}

return user, nil

}

func UserUpdate(c context.Context, id int, req models.User) error {

db := utils.GetDB()

_, err := db.Exec(
	c,
	`UPDATE users
	 SET name=$1,
	     email=$2,
	     password=$3,
	     role=$4
	 WHERE id=$5`,
	req.Name,
	req.Email,
	req.Password,
	req.Role,
	id,
)

return err

}

func UserDelete(c context.Context, id int) error {

db := utils.GetDB()

_, err := db.Exec(
	c,
	`DELETE FROM users WHERE id=$1`,
	id,
)

return err

}
