package user

import (
	"database/sql"
	"fmt"
)

type UserDto struct {
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
