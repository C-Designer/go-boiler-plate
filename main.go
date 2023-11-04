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
	// JSON을 구조체로 변환하기 쉽도록 자동매핑 설정
}

func main() {

	db, _ := config.ConnectionDB()

	http.HandleFunc("/api/v1/user/", func(w http.ResponseWriter, r *http.Request) {
		user.UserController(w, r, db)
	})	// 핸들러 등록

	log.Fatal(http.ListenAndServe(":5050", nil))	// HTTP서버 실패시 실행 중단

}

//go 언어 설치하고 환경 설정 해야됨
