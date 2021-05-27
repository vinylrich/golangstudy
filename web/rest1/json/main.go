package main

import (
	"golangstudy/web/myapp"
	"net/http"
)

//Handler: 그냥 인터페이스 근데 이제 servehttp를 곁들인
func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
