package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func updateEmail(u *User) {
	u.Email = "jo.bob@gmailcom";
}

func main() {
	u1 := User{ID: 1, FirstName: "Jo", LastName: "Bob", Email: "jo.bob@bob.com"}
	fmt.Println(u1)
	updateEmail(&u1)
	fmt.Println(u1)
}
