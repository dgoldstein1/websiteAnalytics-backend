package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "strconv"
)

func main() {
    loggerOn, _ := strconv.ParseBool(os.Getenv("LOGGER"))
    router := NewRouter(loggerOn)

    fmt.Println("Serving on port " + os.Getenv("PORT"))

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
