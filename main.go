package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func DBConfig() *sql.DB {
	db, err := sql.Open("mysql", "root:12341234@tcp(localhost:3306)/go_test")

	if err != nil {
		fmt.Printf("Error %s when opening DB\n", err)
	}
	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Minute * 3)

	query := `CREATE TABLE IF NOT EXISTS User(id int primary key auto_increment, name text,  email varchar(20) ,
		age int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	res, err1 := db.Exec(query)

	if err1 != nil {
		fmt.Printf("query failed %s \n", err1)

	}
	fmt.Printf("query success %+v \n", res)

	return db
}

func main() {

	db := DBConfig()

	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:

			query := `select id,name from User`

			rows, err := db.Query(query)

			if err != nil {
				panic(err.Error())
			}

			var usersArray []Users

			for rows.Next() {
				users := Users{}
				rows.Scan(&users.Id, &users.Name)

				usersArray = append(usersArray, users)
			}

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			json.NewEncoder(w).Encode(usersArray)

			// user.FindUser()

			// board.FindBoard()

		case http.MethodPost:

			println(r.PostFormValue("test"))

			fmt.Printf(" req.body : %+v  ", r.PostForm)

			body := r.PostForm

			query := `insert into User (email, name, age) values (?,?,?)`

			stmp, _ := db.Prepare(query)

			_, err := stmp.Exec(body.Get("email"), body.Get("name"), body.Get("age"))

			if err != nil {
				log.Fatal(err)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			// return
		}
	})

	log.Fatal(http.ListenAndServe(":5050", nil))

}

//go 언어 설치하고 환경 설정 해야됨
