package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

//////////////
// HANDLERS //
//////////////


type Visit struct {
	IpAddress     string `form:"ipAddress" json:"ipAddress" binding:"required"`
	Location string `form:"location" json:"location" binding:"required"`
}

// same as visit, but added timestamp (after insert into DB)
type VisitEntry struct {
	Data Visit
	Timestamp string
}

/**
 * writes list of all site visits
 **/
func getAllVisits(c *gin.Context) {
	values := readAllRows()
	c.String(http.StatusOK, string(values[:]))
}

/**
 * Adds a visit to the postgres DB
 * TODO
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
 * TODO
 **/
func showByIp(c *gin.Context) {
	// ip := c.Param("ip")
	// c.String(http.StatusOK, "Get visit by %s", ip)
}


