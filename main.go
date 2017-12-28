package main

import (
    "fmt"
    "net/http"
    "os"
    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

// what a user uses to POST
type Visit struct {
    IpAddress     string `form:"ipAddress" json:"ipAddress" binding:"required"`
    Location string `form:"location" json:"location" binding:"required"`
}

// same as visit, but added timestamp (after insert into DB)
type VisitEntry struct {
    Data string
    Timestamp string
}

// utilts -- use MAX_INT for nil as int
const (
    NO_INPUT = "9223372036854775807"
    NO_INPUT_INT = 9223372036854775807
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

    http.ListenAndServe(":" + os.Getenv("PORT"), router)
}
