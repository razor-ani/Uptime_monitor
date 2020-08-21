package handler

import (
	"Uptime-monitor/db"
	"Uptime-monitor/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ActivateUrls ..
func ActivateUrls(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)

	if UIDinfo[id].ACTIVATION {
		c.JSON(200, gin.H{
			"Error": "Already inactive",
		})
		return
	}
	db.UpdateActivation(uint64(ID), true)
	go helper.PeriodicCheck(uint64(ID))
}
