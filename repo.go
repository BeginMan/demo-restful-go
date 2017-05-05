package main

import "fmt"

var (
    CurrendId int
    todos Todos
)

func init() {
    createTodo(Todo{Name: "go json"})
    createTodo(Todo{Name: "go http"})
}


func createTodo(t Todo) Todo {
    CurrendId += 1
    t.Id = CurrendId
    todos = append(todos, t)
    return t
}


func FindTodo(id int) Todo {
    for _, todo := range todos {
        if todo.Id == id {
            return todo
        }
    }
    return Todo{}
}


func DestoryTodos(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
