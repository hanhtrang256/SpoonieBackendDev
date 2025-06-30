package main

import (
	"fmt"
	"net/http"
	"os"

	// "os"

	"github.com/hanhtrang256/spoonie-testing/application"
	// "github.com/jackc/pgx/v5"
)

func main() {
	// conn, pgerr := application.ConnectPostgres("postgres", "trang256", "localhost", "5432", "spooniedatabase")
	username := "minhi1"
	password := "eBIFZ8n6FesgQPwUJairg6rJ73Z8wOKD"
	hostname := "dpg-d1fai8qdbo4c739qfsq0-a.singapore-postgres.render.com"
	dbport := "5432"
	dbname := "spooniedatabase"
	conn, pgerr := application.ConnectPostgres(username, password, hostname, dbport, dbname)

	if pgerr != nil {
		fmt.Print("Error connecting database!")
		return
	}

	http.HandleFunc("/", application.DisplayHomePage)
	http.HandleFunc("/login", application.UserLoginAuth(conn))
	http.HandleFunc("/signup", application.UserSignUp(conn))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)
	// http.ListenAndServe("localhost:12001", nil)
	// application.InsertUser(conn, application.Users{Username: "minhi1", Password: "uia256", Role: "admin", Created_at: time.Now()})
}
