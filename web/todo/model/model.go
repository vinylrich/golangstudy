package model

import (
	"time"
)

/* go-sqlite3 패키지는 cgo를 사용하기때문에
c 컴파일러가 있어야 한다.
하지만 window에는 자체 c 컴파일러인
ms-complier가 있긴 하지만
표준을 지키지 않아서 에러가 발생 할 수 있다.
그래서 mingw라는 녀석이 필요하다.
ms-complier를 mingw로 필터링하고,
cgo패키지를 사용 할 수 있다.*/

//Todo struct
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	// handler = newMemoryHandler()
	handler = newSqliteHandler()
}
func GetTodos() []*Todo {
	return handler.getTodos()
}

func AddTodo(name string) *Todo {
	return handler.addTodo(name)
}

func RemoveTodo(id int) bool {
	return handler.removeTodo(id)
}
func CompleteTodo(id int, complete bool) bool {
	return handler.completeTodo(id, complete)
}
