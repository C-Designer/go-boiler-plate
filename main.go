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

type users map[int] User

func (u users) ServeHTTP(w http.ResponseWriter, r *http.Request){
	for nickName,email  := range u {
		fmt.Printf("email : %v nickName : %v",email,nickName)
	}
}

func main() {
 	userss := users{
		1:{Email: "Test",NickName: "sss"},
	}
	pk := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			json.NewEncoder(w).Encode(userss)

		case http.MethodPost:
			var user User
			
			println(r.Body)

			json.NewDecoder(r.Body).Decode(&user)

			println(&user,"user")
			pk =+ pk +1
			userss[pk] = user

			json.NewEncoder(w).Encode(user)
			println(&user,"user")

		}
	})

	log.Fatal(http.ListenAndServe(":5050", userss))

}

//go 환경 셋팅
//go 스타일 낭독
//go arch
