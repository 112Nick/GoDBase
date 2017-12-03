package dao

import (
	m "../models"
	"../database"
	"fmt"
)

func CreateForum(body *m.Forum) error {
	tx := database.Connection
	fmt.Print("asasd")
	//forum := m.Forum{}
	_ , err := tx.Query(`INSERT INTO forum(slug, title , "user", posts, threads) VALUES($1,$2,$3,$4,$5)`,
		body.Slug, body.Title, body.User, body.Posts, body.Threads)//.
			//Scan(&forum.Slug, &forum.Title, &forum.User, &forum.Posts, &forum.Threads)
	if err != nil {
		return err
	}
	//tx.Commit()
	return nil
}

func GetForumBySlug(slug string) *m.Forum {
	tx := database.Connection
	forum := m.Forum{}
	res, _ := tx.Query(`SELECT slug::text, title, "user", posts, threads FROM forum WHERE slug = $1`, slug)//.
	if res.Next() {
		//tx.Rollback()
		res.Scan(&forum.Slug, &forum.Title, &forum.User, &forum.Posts, &forum.Threads)
		fmt.Println(&forum)
		return &forum
	}
	//tx.Commit()
	return nil
}
