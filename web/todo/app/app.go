package app

import (
	"net/http"
	"strconv"

	"golangstudy/web/todo/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}
func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	//3Tierweb FRONT BACK DB

	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list) //web에 반환시켜주기
}

//web 상에 저장->post불러오기
func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") //formvalue가 "name"인 데이터를 받음
	todo := model.AddTodo(name)
	//저쪽 데이터에도 반영을 해줘야함
	rd.JSON(w, http.StatusCreated, todo)

}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //todos 뒤에 있는{id:[0-9]+}를 id에 저장하주는 func
	id, _ := strconv.Atoi(vars["id"])
	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}
func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //todos 뒤에 있는{id:[0-9]+}를 id에 저장하주는 func
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := model.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

//왜 리펙토링 할 때

//web 서버에서 in-memory data를 db로 빼면 3tier web이 됨
func MakeHandler() http.Handler {

	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	return r
}
