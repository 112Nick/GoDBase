package models


type Forum struct {
	Id		int		`json:"id"`
	Posts	int  	`json:"posts"`
	Slug    string 	`json:"slug"`
	Threads int  	`json:"threads"`
	Title   string 	`json:"title"`
	User    string 	`json:"user"`
}


