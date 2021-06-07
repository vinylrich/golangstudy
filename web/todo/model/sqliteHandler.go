package model

import (
	"database/sql"
	"log"
	"time"

	//밑줄 _은 명시적으론 사용하지 않지만
	//암묵적으로 사용한다는 뜻
	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos() []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id,name,completed,createdAT FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}
	return todos
}
func (s *sqliteHandler) AddTodo(name string) *Todo {
	stmt, err := s.db.Prepare("INSERT INTO todos (name,completed,createdAt) VALUES(?,?,datetime('now'))")
	if err != nil {
		log.Panic("stmt err ", err.Error())
	}
	rst, err := stmt.Exec(name, false)
	if err != nil {
		log.Panic("rst err ", err.Error())
	}
	id, _ := rst.LastInsertId()
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()
	return &todo
}
func (s *sqliteHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare("DELETE FROM todos where id=?")
	if err != nil {
		log.Panic("remove stmt err", err.Error())
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		log.Panic("remove rst err", err.Error())
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare("UPDATE todos SET completed=? WHERE id=?")
	if err != nil {
		log.Panic("complete stmt err", err.Error())
	}
	rst, err := stmt.Exec(complete, id)
	if err != nil {
		log.Panic("Complete rst err", err.Error())
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

func newSqliteHandler(filepath string) DBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Panic("make Handler database err ", err.Error())
	}
	statement, err := database.Prepare( //Auto INCREMENT= 자기 혼자서 ++1되는것
		`CREATE TABLE IF NOT EXISTS todos (
			id        INTEGER  PRIMARY KEY AUTOINCREMENT,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		)`)
	if err != nil {
		log.Panic("make handler query err ", err.Error())
	}
	statement.Exec()
	return &sqliteHandler{db: database}
}
