package models

import "time"

type Thread struct {
	Id			int			`json:"id"`
	ForumId		int  		`json:"forumID"`
	Votes  		int 		`json:"votes"`
	Slug		string  	`json:"slug"`
	Author   	string 		`json:"author"`
	Title    	string 		`json:"title"`
	Forum		string 		`json:"forum"`
	Created		time.Time 	`json:"created"`
	Message		string		`json:"message"`
}

