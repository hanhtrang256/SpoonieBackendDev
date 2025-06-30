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
	Signal     string    `json:"signal"`
	Id         string    `json:"id"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"date created"`
}

func ConvertUserJSON(user Users, alert string) UserJSON {
	return UserJSON{Signal: alert, Id: user.Id, Password: user.Password, Username: user.Username, Role: user.Role, Created_at: user.Created_at}
}

func WriteJSON(writer http.ResponseWriter, userJSON UserJSON) {
	writer.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(userJSON)
	if err != nil {
		fmt.Println("Error converting JSON")
		return
	}
	_, err2 := writer.Write(jsonData)
	if err2 != nil {
		fmt.Println("Error writing JSON")
		return
	}
}

type TmpJSON struct {
	Hello string `json:"hello"`
}

func DisplayHomePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	tmp := TmpJSON{Hello: "by minhi1"}
	jsonData, err := json.Marshal(tmp)
	if err != nil {
	}
	_, err2 := writer.Write(jsonData)
	if err2 != nil {
	}
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserLoginAuth(conn *pgx.Conn) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Method", request.Method)
		if request.Method != http.MethodPost {
			http.Error(writer, "Method not allowed!", http.StatusMethodNotAllowed)
			return
		}

		var loginreq AuthRequest
		// read the request body (in json format) and save into loginreq
		err := json.NewDecoder(request.Body).Decode(&loginreq)

		if err != nil {
			http.Error(writer, "JSON decoding error <log in>!", http.StatusBadRequest)
			return
		}

		// username := request.URL.Query().Get("username")
		// password := request.URL.Query().Get("password")

		signal := FindUser(conn, loginreq.Username, loginreq.Password)
		var alert string
		if signal == -1 {
			alert = "Username or password is incorrect!"
		} else if signal == 0 {
			alert = "Wrong password!"
		} else if signal == 1 {
			alert = "Successful login!"
		} else {
			alert = "Bruh what the hell!"
		}

		user := GetUser(conn, loginreq.Username, loginreq.Password)
		userJSON := ConvertUserJSON(user, alert)
		WriteJSON(writer, userJSON)

		// if signal == -1 {
		// 	fmt.Println("Username already exists!")
		// 	// writer.Write([]byte("Username already exists!"))
		// } else if signal == 0 {
		// 	fmt.Println("Wrong password!")
		// 	// writer.Write([]byte("Wrong password!"))
		// } else if signal == 1 {
		// 	fmt.Println("Successful login!")
		// 	// writer.Write([]byte("Success login!"))
		// 	user := GetUser(conn, loginreq.Username, loginreq.Password, scanID)
		// 	userJSON := ConvertJSON(user)
		// 	WriteJSON(writer, userJSON)
		// } else {
		// 	writer.Write([]byte("Bruh what the hell!"))
		// }
	}
}

func UserSignUp(conn *pgx.Conn) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Method", request.Method)
		if request.Method != http.MethodPost {
			http.Error(writer, "Method not allowed!", http.StatusMethodNotAllowed)
			return
		}

		var signupreq AuthRequest
		// read the request body (in json format) and save into loginreq
		err := json.NewDecoder(request.Body).Decode(&signupreq)

		if err != nil {
			http.Error(writer, "JSON decoding error <sign up>!", http.StatusBadRequest)
			return
		}

		user := Users{Username: signupreq.Username, Password: signupreq.Password, Role: "user", Created_at: time.Now()}
		err = InsertUser(conn, user)

		if err != nil {
			http.Error(writer, "Some error inserting user! <user-sign-up>", http.StatusBadRequest)
			return
		}

		userJSON := ConvertUserJSON(user, "successful register")
		writer.Write([]byte("Successfully register!"))
		WriteJSON(writer, userJSON)
	}
}
