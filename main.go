package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	NickName string `json:"nickName"`
	Email string `json:"email"`
}


func test (users map[int]User ){
	fmt.Println(users)
}

func main() {

	 users := make(map[int]*User)

	pk := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			
		case http.MethodGet:
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(users)

		case http.MethodPost:
			var user User
			
			println(r.Body)

			json.NewDecoder(r.Body).Decode(&user)

			pk =+ pk +1
			users[pk] = &user

			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(user)
		}
	})

	log.Fatal(http.ListenAndServe(":5050", nil))

}

//go 환경 셋팅
//go 스타일 낭독
//go arch
