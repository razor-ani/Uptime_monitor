package handler

import (
	"Uptime-monitor/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

//DeactivateUrls ..
func DeactivateUrls(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	varLock.Lock()
	if !(UIDinfo[id].ACTIVATION) {
		c.JSON(200, gin.H{
			"Error": "Already inactive",
		})
		return
	}
	varLock.Unlock()
	db.UpdateActivation(uint64(ID), false)
}
