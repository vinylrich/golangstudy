package main

import (
	"golangstudy/web/logininPage/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Start App")
	http.ListenAndServe(":3000", n)
}
