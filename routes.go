package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

//////////////
// HANDLERS //
//////////////

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

/**
 * writes list of all site visits
 * TODO
 **/
func getAllVisits(c *gin.Context) {
	c.String(http.StatusOK, "All visits")
}

/**
 * Adds a visit to the postgres DB
 * TODO
 **/
func addVisit(c *gin.Context) {
	var person Person
	if c.BindJSON(&person) == nil {
		c.JSON(200, person)
	}
}

/**
 * writes list of all site visits
 * TODO
 **/
func showByIp(c *gin.Context) {
	ip := c.Param("ip")
	c.String(http.StatusOK, "Get visit by %s", ip)
}


