package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type CommonResposne struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}

func Response(w http.ResponseWriter, data interface{}, status int, error error, recoverError ...interface{}) {
	var res CommonResposne

	if status == http.StatusOK {
		res.Data = data
	} else {
		if error != nil {
			res.Error = error.Error()
		} else {
			res.Error = recoverError
		}
	}
	res.Status = status

	w.Header().Set("Content-Type", "application/json;")
	json.NewEncoder(w).Encode(res)
}

func UserController(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r)
			Response(w, nil, http.StatusBadRequest, nil, r)
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
		// 별도 라우팅 패키지를 사용하지않으면 query Param은 직접 URL 구문 분석해야됨

		// string의 zeroValue는 ""
		stringTypeId := strings.TrimPrefix(r.URL.Path, "/api/v1/user/")

		// int의 zeroValue는 0
		conversionNumberTypeId, _ := strconv.Atoi(stringTypeId)

		// param이 없다는건 findAllUser함수를 호출하는것
		if conversionNumberTypeId == 0 {
			result, err := Service.findAllUser()

			if err != nil {
				Response(w, nil, http.StatusBadRequest, err)
				return
			}

			Response(w, result, http.StatusOK, nil)

			// param 으로 넘오면 id로 특정 User 찾기
		} else {

			result, err := Service.findDetailUser(conversionNumberTypeId)

			if err != nil {
				switch err.Error() {
				case "NOT FOUND":
					Response(w, nil, http.StatusBadRequest, errors.New("user not found"))
				default:
					Response(w, nil, http.StatusBadRequest, err)
				}
				return
			}
			Response(w, result, http.StatusOK, nil)
		}

	}

}
