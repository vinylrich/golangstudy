package model

import (
	"log"
	"time"
)

type memoryHandler struct {
	todoMap map[int]*Todo
}

//먼저 분리하고, model.go

func (m *memoryHandler) GetTodos(sessionId string) []*Todo {
	list := []*Todo{} //key=id,value=Todo struct
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddTodo(name string, sessionId string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		log.Println("completed delete")
		return true
	} else {
		return false
	}

}
func (m *memoryHandler) CompleteTodo(id int, complete bool) bool {
	if _, ok := m.todoMap[id]; ok {
		log.Println("Success Check")
		m.todoMap[id].Completed = complete
		return true
	} else {
		return false
	}
}

func (m *memoryHandler) Close() {

}
func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}
