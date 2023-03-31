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

	http.HandleFunc("/api/v1/user", func(w http.ResponseWriter, r *http.Request) {

		user.UserController(w, r, db)
		// switch r.Method {

		// case http.MethodGet:

		// 	query := `select id,name from User`

		// 	rows, err := db.Query(query)

		// 	if err != nil {
		// 		panic(err.Error())
		// 	}

		// 	var usersArray []Users

		// 	for rows.Next() {
		// 		users := Users{}
		// 		rows.Scan(&users.Id, &users.Name)

		// 		usersArray = append(usersArray, users)
		// 	}

		// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 	json.NewEncoder(w).Encode(usersArray)

		// 	// user.FindUser()

		// case http.MethodPost:

		// 	println(r.PostFormValue("test"))

		// 	fmt.Printf(" req.body : %+v  ", r.PostForm)

		// 	body := r.PostForm

		// 	query := `insert into User (email, name, age) values (?,?,?)`

		// 	stmp, _ := db.Prepare(query)

		// 	_, err := stmp.Exec(body.Get("email"), body.Get("name"), body.Get("age"))

		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}

		// 	w.Header().Set("Content-Type", "application/json")
		// 	json.NewEncoder(w).Encode(body)
		// 	// return
		// }
	})

	log.Fatal(http.ListenAndServe(":5050", nil))

}

//go 언어 설치하고 환경 설정 해야됨
