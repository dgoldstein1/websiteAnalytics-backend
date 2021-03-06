package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

// what a user uses to POST
type Visit struct {
	Href         string  `form:"href" json:"href"`
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

	// start background lookup if enabled
	if os.Getenv("BACKGROUND_LOOKUP_ENABLED") == "true" {
		go backgroundLookup()
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

	fmt.Println("Serving on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}

// attempts to lookup empty entries
func backgroundLookup() {
	interval, err := strconv.Atoi(os.Getenv("BACKGROUND_LOOKUP_INTERVAL"))
	if err != nil {
		fmt.Printf("Could not convert BACKGROUND_LOOKUP_INTERVAL: %v", err)
		os.Exit(1)
	}
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	for range ticker.C {
		fmt.Println("looking for empty entires")
		err := updateAllEmptyEntries()
		if err != nil {
			fmt.Printf("error updateAllEmptyEntries(): %v", err)
		}
	}
}
