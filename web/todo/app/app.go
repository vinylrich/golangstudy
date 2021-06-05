package app

import (
	"log"
	"net/http"
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
	log.Println(list)
	rd.JSON(w, http.StatusOK, list)
}
func addTestData() {
	currentTime := time.Now()
	todoMap[1] = &Todo{1, "Milk", true, currentTime}
	todoMap[2] = &Todo{2, "exasdaf", false, currentTime}
	todoMap[3] = &Todo{3, "aasdsaf", false, currentTime}
}
func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	addTestData()

	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/todos", getTodoListHandler)

	return r
}
