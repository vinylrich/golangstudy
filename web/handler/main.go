package main

import (
	"fmt"
	"net/http"
)

//과거의 HTTP는 한 장의 문서를 주고받는
//Client Render html틀을 가지고
//server에서 response받아서 동적으로 랜더링
//대용량 부하를 어떻게 관리하고 넘길것인가

type fooHandler struct{}
type aHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello foo")
}
func (a *aHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hasdasd")
}
func BarHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}
func main() {
	mux := http.NewServeMux() //기존에는 http에 바로 다이렉트로 등록했는데, mux라는 인스턴스를 만들어서 http에 정적으로 등록하는게 아닌 객체를 만듦
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
	mux.HandleFunc("/bar", BarHandler) //경로에 따라서 라우터를 지정"/bar"

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux) //포트번호,라우터

	//1. URL?NAME=TUCKER
	//2. Body

	//string 형태의 dataformat
}
