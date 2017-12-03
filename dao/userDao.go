package dao

import (
	m "../models"
	"../database"
)

func CreateUser(body *m.User) error {
	tx := database.Connection
	_, err := tx.Query(`INSERT INTO users(fullname, nickname, email, about)
									VALUES($1,$2,$3,$4)`, body.Fullname, body.Nickname, body.Email, body.About)
	if err != nil {
		//tx.Rollback()
		return err
	}
	//tx.Commit()
	return nil
}


func GetUserByNickname(nickname string) *m.User {
	tx := database.Connection
	user := m.User{}
	res, _ := tx.Query(`SELECT nickname::text, fullname, email::text, about FROM users
									WHERE nickname = $1`, nickname)
	if res.Next() {
		res.Scan(&user.Nickname, &user.Fullname, &user.Email, &user.About)
		//tx.Rollback()
		return &user
	}
	//tx.Commit()
	return nil
}

func GetUserByNickOrEmail(nickname, email string, users *[]m.User) error{
	tx := database.Connection

	exists, err := tx.Query(`SELECT nickname::text, fullname, email::text, about FROM users
									WHERE nickname = $1 OR email = $2`, nickname , email)
	if err !=nil {
		return err
	}
	for exists.Next() {
		user := m.User{}
		exists.Scan(&user.Nickname, &user.Fullname, &user.Email, &user.About)
		*users = append(*users, user)
	}
	return nil
}

