package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dudckd6744/go-sever-study/modules/board"
	"github.com/dudckd6744/go-sever-study/modules/user"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	NickName string `json:"nickName"`
	Email    string `json:"email"`
}


func DBConfig() {
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
}

func main() {

	DBConfig()

	users := make(map[int]*User)

	pk := 0
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

			case http.MethodGet:
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(users)
				user.FindUser()
				board.FindBoard()
			case http.MethodPost:
				// var user User

				println(r.PostFormValue("test"))
				// println(r.Body)

				fmt.Printf(" req.body : %+v  ",r.PostForm)

				body := r.PostForm

				fmt.Print(body)
				// json.NewDecoder(r.Body).Decode(&user)
				pk = +pk + 1
				// var ee User = User(body)
				// users[pk] = User(body)

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(body)
				// return
			}
	})

	log.Fatal(http.ListenAndServe(":5050", nil))

}

//go 환경 셋팅
//go 스타일 낭독
//go arch
