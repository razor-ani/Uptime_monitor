package handler

import (
	"Uptime-monitor/db"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

//DeleteUrls ..
func DeleteUrls(c *gin.Context) {
	log.Printf("Accessing request params...")
	id := c.Param("id")

	err := db.DeleteWithID(id)
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "message": fmt.Sprintf("%v", err)})
		return
	}
}
