package user

import (
	"database/sql"
	"fmt"
)

type UserDto struct {
	Email string
	Name  string
}

type UserRaw struct {
	Id    int
	Email string
	Name  string
}

type UserRepository struct {
	DB *sql.DB
}

var Repository UserRepository

func (r *UserRepository) AssignDB(db *sql.DB) {
	r.DB = db
}

func (r *UserRepository) createUser(payload UserDto) (sql.Result, error) {

	query := `insert into User (email, name) values (? , ?)`

	fmt.Print("test", payload)
	result, err := r.DB.Exec(query, payload.Email, payload.Name)

	fmt.Print("test", result)
	if err != nil {
		fmt.Print("test", err)
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) findAllUser() ([]UserRaw, error) {
	var raws []UserRaw

	query := `select id,email,name from User`

	result, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for result.Next() {
		var raw UserRaw

		err := result.Scan(&raw.Id, &raw.Email, &raw.Name)

		if err != nil {
			return nil, err
		}

		raws = append(raws, raw)
	}

	fmt.Println(&raws)
	return raws, nil
}
