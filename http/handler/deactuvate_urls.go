package handler

import (
	"Uptime-monitor/db"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

//DeactivateUrls ..
func DeactivateUrls(c *gin.Context) {
	log.Printf("Accessing request params...")
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
	err := db.UpdateActivation(uint64(ID), false)
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "message": fmt.Sprintf("%v", err)})
		return
	}
}
