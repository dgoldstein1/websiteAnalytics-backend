package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

func main() {
    fmt.Println("Serving on port " + os.Getenv("PORT"))

    router := gin.Default()
    router.Use(gin.Logger())

    // declare routes
    router.GET("/", getAllVisits)
    router.GET("/visits", getAllVisits)
    router.POST("visits", addVisit)
    router.GET("/visits/ip/:ip",showByIp)

    newRepo("bolt.db")

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), router))
}
