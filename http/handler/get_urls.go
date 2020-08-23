package handler

import (
	"Uptime-monitor/db"
	"Uptime-monitor/helper"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUrls ...
func GetUrls(c *gin.Context) {
	log.Printf("Accessing request params...")
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(422, gin.H{"code": "422", "message": "Unprocessable Entity"})
		return
	}

	ID := uint64(i)

	data, err := db.FetchAllData(ID)
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "message": fmt.Sprintf("%v", err)})
		return
	}
	log.Printf("Creating response object...")
	c.JSON(200, helper.CreateResponse(data))
}
