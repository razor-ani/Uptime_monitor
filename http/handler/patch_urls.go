package handler

import (
	"Uptime-monitor/db"
	"Uptime-monitor/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

//PatchUrls ...
func PatchUrls(c *gin.Context) {

	data := db.EntryData{}

	data.Crawltimeout, _ = strconv.Atoi(c.PostForm("crawl_timeout"))
	data.Frequency, _ = strconv.Atoi(c.PostForm("frequency"))
	data.Failurethreshold, _ = strconv.Atoi(c.PostForm("failure_threshold"))
	data.UUID = c.Param("id")
	varLock.Lock()
	data.ID = UIDinfo[data.UUID].ID
	data.Status = UIDinfo[data.UUID].STATUS
	varLock.Unlock()
	db.UpdateEntry(&data)

	c.JSON(200, helper.CreateResponse(&data))
}
