package models

import (
	"errors"

	"github.com/abhijitpattar/gin-rest-go/db"
	"github.com/abhijitpattar/gin-rest-go/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT into USERS(email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashpass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashpass)
	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()

	u.ID = userID
	return err
}

func (u *User) ValidateCredentails() error {
	query := `
	select id, password from users
	where email = ?
	`
	//fmt.Print("query = ", query)
	//fmt.Println("email =", u.Email)
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Credentials Invalid")
	}

	//fmt.Println("retrievedPassword = ", retrievedPassword)
	//fmt.Println("u.Password =", u.Password)
	passwordIsValid := utils.CheckPasswordhash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Credentials Invalid")
	}

	return nil
}
