package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type CommonResposne struct {
	Data   interface{} `json: "data"`
	Status int         `json: "status"`
	Error  interface{} `json: "error"`
}

func Response(w http.ResponseWriter, data interface{}, status int, error interface{}) {
	var res CommonResposne

	if status == http.StatusOK {
		res.Data = data
	} else {
		res.Error = error
	}
	res.Status = status

	w.Header().Set("Content-Type", "application/json;")
	json.NewEncoder(w).Encode(res)
}

func UserController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer func() {
		if r := recover(); r != nil {
			Response(w, nil, http.StatusBadRequest, r)
		}
	}()

	Service.InitService(db)

	switch r.Method {
	case http.MethodPost:
		var body struct {
			Email string
			Name  string
		}

		json.NewDecoder(r.Body).Decode(&body)

		_, err := Service.createUser(body)

		if err != nil {
			Response(w, nil, http.StatusBadRequest, err)
		}

		Response(w, "OK", http.StatusOK, nil)
	case http.MethodGet:
		result, err := Service.findAllUser()

		if err != nil {
			Response(w, nil, http.StatusBadRequest, err)
		}

		Response(w, result, http.StatusOK, nil)
	}

}
