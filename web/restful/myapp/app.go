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

//실무에서는 update용 struct를 따로 만듦
//실제에서는 lastname을 없애거나 할 때
//default값과 ""값을 구분 할 수 없기때문에
//flag를 달아줌으로서 바꿀건지 안 바꿀건지 구분
type UpdateUser struct {
	ID               int `json:"id"`
	UpdatedFirstName bool
	FirstName        string `json:"first_name"`
	UpdatedLastName  bool
	LastName         string `json:"last_name"`
	UpdatedEmailName bool
	Email            string `json:"email"`

	CreateAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
func usersHandler(w http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users to Print")
		return
	}
	users := []*User{}

	for _, u := range userMap { //value값(구조체)을 넣음
		users = append(users, u)
		fmt.Println(u)
	}

	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
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

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	_, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Delete User ID:", id)
		return
	}
	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted ID:", id)
}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	updateUser := new(User)
	/*
		{
			"id":%d,
			"first_name":"update"
		}
	*/
	err := json.NewDecoder(r.Body).Decode(updateUser)
	//decode:바디에서 json데이터 받아오기
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	//1번 id를 가지면
	//1번 id를 가진 user 구조체를 가져옴
	user, ok := userMap[updateUser.ID]
	/*
		{
			"first_name":"junwoo",
			"last_name":"kim",
			"email":"whktjd0109@gmail.com"
		}
	*/
	if !ok {
		updateUser.ID = 1
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User ID:", updateUser.ID)
		return
	}
	//실제로 client가 수정하고 싶은
	//부분만 수정 가능하게 만들기
	if updateUser.FirstName == "" {
		updateUser.FirstName = user.FirstName
	} //
	if updateUser.LastName == "" {
		updateUser.LastName = user.LastName
	}
	if updateUser.Email == "" {
		updateUser.Email = user.Email
	}
	data, _ := json.Marshal(updateUser) //
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
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
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", getUsersInfoHandler).Methods("GET")

	// mux.HandleFunc("/users", deleteUserHandler).Methods("DELETE")
	return mux
}
