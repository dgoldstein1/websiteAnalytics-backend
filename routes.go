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

/**
 * writes list of all site visits
 * TODO
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
	// var json Visit
	// if err := c.ShouldBindJSON(&json); err == nil {
	// 	c.JSON(200, json)
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }
}

/**
 * writes list of all site visits
 * TODO
 **/
func showByIp(c *gin.Context) {
	// ip := c.Param("ip")
	// c.String(http.StatusOK, "Get visit by %s", ip)
}


