package app

import (
	"log"
	"net/http"
	"strconv"

	"golangstudy/web/todo/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	//3Tierweb FRONT BACK DB

	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list) //web에 반환시켜주기
}

//web 상에 저장->post불러오기
func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") //formvalue가 "name"인 데이터를 받음
	todo := a.db.AddTodo(name)
	//저쪽 데이터에도 반영을 해줘야함
	rd.JSON(w, http.StatusCreated, todo)

}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //todos 뒤에 있는{id:[0-9]+}를 id에 저장해주는 func
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}
func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //todos 뒤에 있는{id:[0-9]+}를 id에 저장해주는 func
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	log.Println(complete)
	ok := a.db.CompleteTodo(id, complete) //false
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

//왜 리펙토링 할 때

func (a *AppHandler) Close() {
	a.db.Close()
}

//web 서버에서 in-memory data를 db로 빼면 3tier web이 됨
func MakeHandler(filepath string) AppHandler {
	r := mux.NewRouter()
	a := AppHandler{
		Handler: r,
		db:      model.NewDBHandler(filepath),
	}
	r.HandleFunc("/", a.indexHandler)
	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	return a
}
