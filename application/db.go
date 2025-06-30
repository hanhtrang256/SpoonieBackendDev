package application

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	Id         string
	Username   string
	Password   string
	Role       string
	Created_at time.Time
}

// Establish connection
func ConnectPostgres(username, password, host, port, dbName string) (*pgx.Conn, error) {
	// Database url format: postgres://username:password@localhost:5432/database_name
	var connString string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	conn, err := pgx.Connect(context.Background(), connString)
	// fmt.Println(conn.Config().Database, conn.Config().Host, conn.Config().Password, conn.Config().User)
	if err != nil {
		fmt.Print("Error connecting with database")
		return nil, err
	}
	// fmt.Println("Pass?")
	return conn, err
}

// Insert new user
// TODO: encrypt user password
func InsertUser(conn *pgx.Conn, user Users) error {
	// Define the SQL insert query
	query := `INSERT INTO Users (username, password, role, created_at) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING id`

	// Execute and query the id with named arguments
	row := conn.QueryRow(context.Background(), query, user.Username, user.Password, user.Role, user.Created_at)
	err := row.Scan(&user.Id)

	if err != nil {
		fmt.Println("Error inserting user with id ", user.Id)
		return err
	}
	fmt.Printf("Successfully add user with id %s!\n", user.Id)
	return nil
}

// Find and return user ID with username and password
func FindUser(conn *pgx.Conn, username string, password string) (int, string) {
	query := `SELECT id,password FROM Users WHERE username = $1`

	row := conn.QueryRow(context.Background(), query, username)
	var scanID, scanPassword string
	err := row.Scan(&scanID, &scanPassword)

	if err != nil {
		return -1, ""
	}

	if scanPassword != password {
		return 0, ""
	}
	return 1, scanID
}

func GetUser(conn *pgx.Conn, username string, password string, id string) Users {
	query := `SELECT * FROM Users WHERE username = $1 AND password = $2 AND id = $3`

	row := conn.QueryRow(context.Background(), query, username, password, id)
	var user Users
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.Created_at)

	if err != nil {
		fmt.Println("Some error getting user!!!")
		return Users{}
	}

	return user
}
