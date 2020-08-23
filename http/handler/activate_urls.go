package handler

import (
	"Uptime-monitor/db"
	"Uptime-monitor/helper"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ActivateUrls ..
func ActivateUrls(c *gin.Context) {
	log.Printf("Accessing request params...")
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)

	if UIDinfo[id].ACTIVATION {
		c.JSON(200, gin.H{
			"Error": "Already active",
		})
		return
	}
	err := db.UpdateActivation(uint64(ID), true)
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "message": fmt.Sprintf("%v", err)})
		return
	}
	go helper.PeriodicCheck(uint64(ID))

}
