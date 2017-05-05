package main

import (
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)


func Index(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "TODO LIST!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }

}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    var todoId int
    var err error
    if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
        panic(err)
    }

    todo := FindTodo(todoId)
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")

    if todo.Id > 0 {
        if err := json.NewEncoder(w).Encode(todo); err != nil {
            panic(err)
        }
        return

    }

    // 404
    w.WriteHeader(http.StatusNotFound)
    if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
        panic(err)
    }
}


func CreateTodo(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    // body, err := ioutil.ReadAll(io.LimitedReader(r.Body, 120000))
    if err != nil {
        panic(err)
    }

    // TODO: close r.Body
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-Type", "application/json;charset=UTF-8")
        w.WriteHeader(422)
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := createTodo(todo)
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }

}
