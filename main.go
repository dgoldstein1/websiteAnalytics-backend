package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// what a user uses to POST
type Visit struct {
	Ip           string  `form:"ip" json:"ip" binding:"required"`
	City         string  `form:"city" json:"city"`
	Country_Code string  `form:"country_code" json:"country_code"`
	Country_Name string  `form:"country_name" json:"country_name"`
	Latitude     float64 `form:"latitude" json:"latitude"`
	Longitude    float64 `form:"longitude" json:"longitude"`
	Metro_Code   int     `form:"metro_code" json:"metro_code"`
	Region_Code  string  `form:"region_code" json:"region_code"`
	Time_Zone    string  `form:"time_zone" json:"time_zone"`
	Zip_Code     string  `form:"zip_code" json:"zip_code"`
	// added when inserted into db, passing this in POST does not do anything
	Visit_Date time.Time `form:"visit_date" json:"visit_date"` // RFC3339 date string
}

// utilts -- use MAX_INT for nil as int
const (
	NO_INPUT       = "9223372036854775807"
	NO_INPUT_INT   = 9223372036854775807
	NO_INPUT_FLOAT = 9223372036854775807.0
)

func main() {
	// connect to db
	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		fmt.Println("no environment variable DATABASE_URL provided, exiting.")
		os.Exit(1)
	}
	connectionSuccess := connectToDb(uri)
	if connectionSuccess == false {
		fmt.Println("Could not connect to db, exiting")
		os.Exit(1)
	}

	// set up routing
	router := gin.Default()
	router.Use(gin.Logger())

	// set base page as readme html
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// initialize REST routes
	router.GET("/visits", getAllVisits)
	router.POST("visits", addVisit)
	router.GET("/visits/ip/:ip", showByIp)

	fmt.Println("Serving on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
