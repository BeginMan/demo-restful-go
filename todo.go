package main

import (
    "time"
)

type Todo struct {
    Id  int
    Name string
    Completed bool
    Due time.Time `json:due`
}

type Todos []Todo


