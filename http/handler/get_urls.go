package handler

import (
	"Uptime-monitor/db"
	"Uptime-monitor/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUrls ...
func GetUrls(c *gin.Context) {

	id := c.Param("id")

	i, _ := strconv.Atoi(id)

	ID := uint64(i)

	data := db.FetchAllData(ID)

	c.JSON(200, helper.CreateResponse(data))
}
