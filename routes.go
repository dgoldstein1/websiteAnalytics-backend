package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

//////////////
// HANDLERS //
//////////////

type Visit struct {
	ip      string `json:"ip"`
	location string    `json:"location"`
}

/**
 * writes list of all site visits
 * TODO
 **/
func getAllVisits(c *gin.Context) {
	values := getValues()
	c.String(http.StatusOK, string(values[:]))

}

/**
 * Adds a visit to the postgres DB
 * TODO
 **/
func addVisit(c *gin.Context) {
	var visit Visit
	if c.BindJSON(&visit) == nil {
		c.JSON(http.StatusOK, visit)
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


