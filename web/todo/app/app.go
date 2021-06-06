package app

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo
var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}
func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	//3Tierweb FRONT BACK DB
	list := []*Todo{} //key=id,value=Todo struct
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list) //web에 반환시켜주기
}

//web 상에 저장->post불러오기
func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name") //formvalue가 "name"인 데이터를 받음
	id := len(todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo

	//저쪽 데이터에도 반영을 해줘야함
	rd.JSON(w, http.StatusOK, todo)

}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //todos 뒤에 있는{id:[0-9]+}를 id에 저장하주는 func
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		log.Println("completed delete")
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}
func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //todos 뒤에 있는{id:[0-9]+}를 id에 저장하주는 func
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
	}
	complete := r.FormValue("complete") == "true"
	if _, ok := todoMap[id]; ok {
		log.Println("Success Check")
		todoMap[id].Completed = complete
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

//web 서버에서 in-memory data를 db로 빼면 3tier web이 됨
func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)

	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("PUT")
	return r
}
