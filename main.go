package main

import (
	"log"
	"net/http"

	"github.com/dudckd6744/go-sever-study/config"
	"github.com/dudckd6744/go-sever-study/modules/user"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {

	db, _ := config.ConnectionDB()

	http.HandleFunc("/api/v1/user/", func(w http.ResponseWriter, r *http.Request) {
		user.UserController(w, r, db)
	})

	log.Fatal(http.ListenAndServe(":5050", nil))

}

//go 언어 설치하고 환경 설정 해야됨
