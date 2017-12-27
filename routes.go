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


/**
 * writes list of all site visits
 **/
func getAllVisits(c *gin.Context) {
	// query string parameters. Default is last month
	from := c.DefaultQuery("from", NO_INPUT)
	to := c.DefaultQuery("to", NO_INPUT)
	ip := c.DefaultQuery("ip", NO_INPUT)
	// cast to ints
	fromInt, fromErr := strconv.Atoi(from)
	toInt, toErr := strconv.Atoi(to)
	// validate inputs
	if (fromErr != nil || toErr != nil) {
		c.String(http.StatusBadRequest, "Bad input " + to + " " + from)
	} else if (fromInt > toInt) {
		c.String(http.StatusBadRequest, "The from date cannot be greater that the to date: input : " + to + " - " + from)
	} else { // input is good, read data
		values, err := readAllRows(ip, toInt, fromInt)
		if (err == nil ) {
			c.String(http.StatusOK, string(values[:]))
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


