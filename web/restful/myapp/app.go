package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var userMap map[int]*User
var lastID int

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`

	CreateAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get userInfo by /users/{id}")
}
func getUsersInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Contents-Type", "application/json")
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

//createUser handler
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) //data받아오기
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request", err)
		return
	}

	user.CreateAt = time.Now()

	lastID++
	user.ID = lastID
	user.CreateAt = time.Now()
	userMap[user.ID] = user
	data, _ := json.Marshal(user) //
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET") //users가 없으면 상위 핸들러가 호출됨
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", getUsersInfoHandler)
	// mux.HandleFunc("/users", deleteUserHandler).Methods("DELETE")
	return mux
}

