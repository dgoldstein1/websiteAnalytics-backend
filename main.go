package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
)

func main() {
    router := NewRouter(true)
    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "hello, world")
}