package models_users

import (
	"context"
	"fmt"
	"log"

	"github.com/abhinavramkumar/go-rss/database"
	"github.com/abhinavramkumar/go-rss/structs"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

var db = database.GetDB()

func CheckIfUserExists(email string) (string, error) {
	var userID string
	query := `select user_id from users where email=$1`
	err := db.QueryRow(context.Background(), query, email).Scan(&userID)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return userID, nil
}

func VerifyPassword(password string, userID string) (bool, error) {
	var storedPwd string
	query := `select password from users where user_id=$1`
	err := db.QueryRow(context.Background(), query, userID).Scan(&storedPwd)
	if err != nil {
		log.Println(err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPwd), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func InsertUser(user structs.UserStruct) (structs.UserStruct, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return structs.UserStruct{}, err
	}
	query := `insert into users (name, email, password) VALUES ($1, $2, $3)`
	_, err = db.Exec(context.Background(), query, user.Name, user.Email, string(hashedPassword))
	if err != nil {
		return structs.UserStruct{}, err
	}

	fmt.Printf("\nUserID: %s", user.UserID)
	var result structs.UserStruct
	query = `select user_id, name, email, created_at, updated_at from users where email=$1`
	rows, _ := db.Query(context.Background(), query, user.Email)
	result, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[structs.UserStruct])

	if err != nil {
		return structs.UserStruct{}, err
	}

	fmt.Printf("Result: %+v", result)

	return result, nil
}
