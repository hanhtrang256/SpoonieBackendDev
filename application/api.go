package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	// "github.com/labstack/echo/v4"
)

type UserJSON struct {
	Id         string    `json:"id"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"date created"`
}

type TmpJSON struct {
	Hello string `json:"hello"`
}

func DisplayHomePage(writer http.ResponseWriter, request *http.Request) {
	// users := []UserJSON{{Id: "6", Username: "minhi1", Role: "Admin", Created_at: time.Now()},
	// 	{Id: "25", Username: "TR", Role: "Admin", Created_at: time.Now()}}

	// jsonData, err := json.Marshal(users)
	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	writer.Header().Set("Content-Type", "application/json")

	// _, err = writer.Write(jsonData)
	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	tmp := TmpJSON{Hello: "by minhi1"}
	jsonData, err := json.Marshal(tmp)
	if err != nil {

	}
	_, err2 := writer.Write(jsonData)
	if err2 != nil {

	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserAuthentication(conn *pgx.Conn) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Method", request.Method)
		if request.Method != http.MethodPost {
			http.Error(writer, "Method not allowed!", http.StatusMethodNotAllowed)
			return
		}

		var loginreq LoginRequest
		// read the request body (in json format) and save into loginreq
		err := json.NewDecoder(request.Body).Decode(&loginreq)

		if err != nil {
			http.Error(writer, "JSON decoding error!", http.StatusBadRequest)
			return
		}

		// username := request.URL.Query().Get("username")
		// password := request.URL.Query().Get("password")

		signal := FindUser(conn, loginreq.Username, loginreq.Password)
		// signal := FindUser(conn, username, password)
		// fmt.Println("username", username)
		// fmt.Println("password", password)
		if signal == -1 {
			writer.Write([]byte("Non-exist"))
		} else if signal == 0 {
			writer.Write([]byte("Wrong password"))
		} else if signal == 1 {
			writer.Write([]byte("Success login"))
		} else {
			writer.Write([]byte("Bruh what the hell"))
		}
	}
}
