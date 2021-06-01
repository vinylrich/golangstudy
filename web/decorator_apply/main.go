package main

import (
	"golangstudy/web/decorator_apply/myapp"
	"log"
	"net/http"
	"time"

	decohandler "golangstudy/web/decoHandler"
)

func logger1(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER1] Started")

	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds())
}
func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER2] Started")

	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed time:", time.Since(start).Milliseconds())
}
func NewHandler() http.Handler {
	h := myapp.NewHandler()
	mux := myapp.NewHandler()
	h = decohandler.NewDecoHandler(mux, logger1)
	h = decohandler.NewDecoHandler(mux, logger2)

	return h
}
func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
