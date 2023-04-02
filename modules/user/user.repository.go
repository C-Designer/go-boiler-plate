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

func (r *UserRepository) findAllUser() (*[]UserRaw, error) {
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

	/**
	struct 타입을 포인터로 return 시키는 이유
	- 데이터를 포인터로 지정안하고 반환할 경우 변수 대입과 함수 인수 전달은 항상 값을 복사하기 때문에 메모리 비효율과 성능 문제를 발생시킨다.
	**/
	return &raws, nil
}
