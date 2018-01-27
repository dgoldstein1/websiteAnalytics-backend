// routes.go

package main

/**
 * Created by David Goldstein 12/2017
 * route handlers
 **/

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "strconv"
)

//////////////
// HANDLERS //
//////////////

/**
 * writes list of all site visits
 **/
func getAllVisits(c *gin.Context) {
	// property filters
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

	// cast values
	fromInt, fromErr := strconv.Atoi(from)
	toInt, toErr := strconv.Atoi(to)
	latitudeFloat, latitudeErr :=  strconv.ParseFloat(latitude, 64)
	longitudeFloat, longitudeErr := strconv.ParseFloat(longitude, 64)
	metroCodeInt, metroCodeErr := strconv.Atoi(metro_code)

	// validate inputs
	if (latitudeErr != nil || longitudeErr != nil || metroCodeErr != nil) {
		c.String(http.StatusBadRequest, "Error parsing latitude, longitude, or metro_code")
	} else if (fromErr != nil || toErr != nil) {
		c.String(http.StatusBadRequest, "Bad input " + to + " " + from)
	} else if (fromInt > toInt) {
		c.String(http.StatusBadRequest, "The from date cannot be greater that the to date: input : " + to + " - " + from)
	} else { // input is good, read data
		// convert params into query object
		query, queryErr := createQueryFromFilters(ip, city, country_code, country_name, latitudeFloat, longitudeFloat, metroCodeInt, region_code, time_zone, zip_code)
		if (queryErr == nil) { // query succesfully created
			// fetch from mongo and write results
			values, err := readAllRows(query, toInt, fromInt)
			if (err == nil) {
				c.String(http.StatusOK, string(values[:]))
			} else {
				c.String(500, err.Error())
			}
		} else { // error creating the query
			c.String(500, queryErr.Error())
		}
	}
}

/**
 * Add a visit to the db
 **/
func addVisit(c *gin.Context) {
	var in Visit
	if err := c.ShouldBindJSON(&in); err == nil {
		out, insertErr := insertRow(in)
		if (insertErr!=nil) {
				c.String(500, insertErr.Error())
			} else {
				c.JSON(200, out)
			}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

/**
 * writes list of all site visits
 **/
func showByIp(c *gin.Context) {
	ip := c.Param("ip")
	values, err := readByIp(ip)
	if (err != nil) {
		c.String(500, err.Error())
	} else {
		c.String(http.StatusOK, string(values[:]))
	}
}


