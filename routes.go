package main

import (
    "net/http"

    "github.com/gorilla/mux"
)


type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandleFunc  http.HandlerFunc
}

type Routes []Route


var routes = Routes{
    Route{"Index", "GET", "/", Index},
    Route{"TodoIndex", "GET", "/todos", TodoIndex},
    Route{"TodoShow", "GET", "/todos/{todoId}", TodoShow},
    Route{"CreateShow", "POST", "/todos", CreateTodo},
}


func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var Handler http.Handler
        // add log
        Handler = Logger(Handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandleFunc)
    }
    return router
}

