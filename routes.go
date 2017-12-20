package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

//////////////
// HANDLERS //
//////////////

type Food struct {
	Name      string `json:"name"`
	Id        int    `json:"id"`
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
	var food Food
	if c.BindJSON(&food) == nil {
		c.JSON(http.StatusOK, food)
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


