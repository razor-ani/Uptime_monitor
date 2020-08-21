package handler

import (
	"Uptime-monitor/db"

	"github.com/gin-gonic/gin"
)

//DeleteUrls ..
func DeleteUrls(c *gin.Context) {
	id := c.Param("id")

	db.DeleteWithID(id)
}
