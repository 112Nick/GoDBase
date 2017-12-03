package models

import "time"

type Post struct {
	Id			int			`json:"id"`
	Parent		int  		`json:"parent"`
	Thread  	int 		`json:"thread"`
	IsEdited	bool  		`json:"isEdited"`
	Author   	string 		`json:"author"`
	Message    	string 		`json:"message"`
	Forum		string 		`json:"forum"`
	Created		time.Time 	`json:"created"`
	Path		[]int		`json:"path"`
}