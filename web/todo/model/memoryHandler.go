package model

import (
	"log"
	"time"
)

type memoryHandler struct {
	todoMap map[int]*Todo
}

type dbHandler interface {
	getTodos() []*Todo
	addTodo(name string) *Todo
	removeTodo(id int) bool
	completeTodo(id int, complete bool) bool
}

//먼저 분리하고, model.go

func (m *memoryHandler) getTodos() []*Todo {
	list := []*Todo{} //key=id,value=Todo struct
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) addTodo(name string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler) removeTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		log.Println("completed delete")
		return true
	} else {
		return false
	}

}
func (m *memoryHandler) completeTodo(id int, complete bool) bool {
	if _, ok := m.todoMap[id]; ok {
		log.Println("Success Check")
		m.todoMap[id].Completed = complete
		return true
	} else {
		return false
	}
}

var handler dbHandler

func newMemoryHandler() dbHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}
