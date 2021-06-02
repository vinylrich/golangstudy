package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}
func main() {
	user1 := User{Name: "junwoo", Email: "whktjd0109@gmail.com", Age: 19}
	user2 := User{Name: "213", Email: "wwasd", Age: 40}
	users := []User{user1, user2}
	tmpl, err := template.New("Tmpl1").ParseFiles("templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
