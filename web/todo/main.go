package main

import (
	"golangstudy/web/todo/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Start App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
