package models

type PostFull struct {
	Post	Post	`json:"post"`
	Author	User	`json:"author"`
	Forum  	Forum	`json:"forum"`
	Thread	Thread 	`json:"thread"`

}
