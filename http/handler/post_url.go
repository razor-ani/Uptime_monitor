package handler

import (
	"Uptime-monitor/db"
	"Uptime-monitor/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

//URLToID is ...

//PostUrls is ..
func PostUrls(c *gin.Context) {
	data := db.EntryData{}

	data.UUID = helper.GetUID()
	data.URL = c.PostForm("url")
	data.Crawltimeout, _ = strconv.Atoi(c.PostForm("crawl_timeout"))
	data.Frequency, _ = strconv.Atoi(c.PostForm("frequency"))
	data.Failurethreshold, _ = strconv.Atoi(c.PostForm("failure_threshold"))

	if helper.Check(data.URL) {
		data.Status = "active"
	} else {
		data.Status = "inactive"
	}
	data.Activate = true

	data.ID = db.InsertURL(&data)

	c.JSON(200, helper.CreateResponse(&data))

	varLock.Lock()
	UIDinfo[data.UUID] = INFO{data.Status, data.ID, data.Activate}
	varLock.Unlock()
	go helper.PeriodicCheck(data.ID)
}
