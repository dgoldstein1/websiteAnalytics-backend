// routes.go

package main

/**
 * Created by David Goldstein 12/2017
 * route handlers
 **/

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//////////////
// HANDLERS //
//////////////

/**
 * writes list of all site visits
 **/
func getAllVisits(c *gin.Context) {
	// property filters
	href := c.DefaultQuery("href", NO_INPUT)
	ip := c.DefaultQuery("ip", NO_INPUT)
	city := c.DefaultQuery("city", NO_INPUT)
	country_code := c.DefaultQuery("country_code", NO_INPUT)
	country_name := c.DefaultQuery("country_name", NO_INPUT)
	latitude := c.DefaultQuery("latitude", NO_INPUT)
	longitude := c.DefaultQuery("longitude", NO_INPUT)
	metro_code := c.DefaultQuery("metro_code", NO_INPUT)
	region_code := c.DefaultQuery("region_code", NO_INPUT)
	time_zone := c.DefaultQuery("time_zone", NO_INPUT)
	zip_code := c.DefaultQuery("zip_code", NO_INPUT)

	// time filters
	from := c.DefaultQuery("from", NO_INPUT)
	to := c.DefaultQuery("to", NO_INPUT)

	// filter type
	query_type := c.DefaultQuery("query_type", "and")

	// cast values
	fromInt, fromErr := strconv.Atoi(from)
	toInt, toErr := strconv.Atoi(to)
	latitudeFloat, latitudeErr := strconv.ParseFloat(latitude, 64)
	longitudeFloat, longitudeErr := strconv.ParseFloat(longitude, 64)
	metroCodeInt, metroCodeErr := strconv.Atoi(metro_code)

	// validate inputs
	if query_type != "and" && query_type != "or" && query_type != "nor" {
		c.String(http.StatusBadRequest, "Cannot filter on "+query_type+". Must be 'and' || 'or' || 'nor.")
	} else if latitudeErr != nil || longitudeErr != nil || metroCodeErr != nil {
		c.String(http.StatusBadRequest, "Error parsing latitude, longitude, or metro_code")
	} else if fromErr != nil || toErr != nil {
		c.String(http.StatusBadRequest, "Bad input "+to+" "+from)
	} else if fromInt > toInt {
		c.String(http.StatusBadRequest, "The from date cannot be greater that the to date: input : "+to+" - "+from)
	} else { // input is good, read data
		// convert params into query object
		visitFilters := Visit{href, ip, city, country_code, country_name, latitudeFloat, longitudeFloat, metroCodeInt, region_code, time_zone, zip_code, time.Now()}
		// fetch from mongo and write results
		values, err := readAllRows(visitFilters, toInt, fromInt, query_type)
		if err == nil {
			c.JSON(http.StatusOK, values)
		} else {
			c.String(500, err.Error())
		}
	}
}

/**
 * Add a visit to the db
 **/
func addVisit(c *gin.Context) {
	var in Visit
	if err := c.ShouldBindJSON(&in); err == nil {
		visit, insertErr := insertRow(in)
		if insertErr != nil {
			c.String(500, fmt.Sprintf("Could not insert doc: %v", insertErr.Error()))
		} else {
			c.JSON(200, gin.H{
				"visit":     visit,
				"totalDocs": docCount(),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
