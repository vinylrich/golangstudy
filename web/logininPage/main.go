package main

import (
	"golangstudy/web/logininPage/app"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()

	log.Println("Start App")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}

}
