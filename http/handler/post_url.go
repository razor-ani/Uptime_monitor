package handler

import (
	"Uptime-monitor/db"
	"Uptime-monitor/helper"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

//URLToID is ...

//PostUrls is ..
func PostUrls(c *gin.Context) {
	data := db.EntryData{}
	log.Printf("Accessing request params...")
	data.UUID = helper.GetUID()
	data.URL = c.PostForm("url")
	var err error = nil

	data.Crawltimeout, err = strconv.Atoi(c.PostForm("crawl_timeout"))
	if err != nil {
		c.JSON(422, gin.H{"code": "422", "message": "Unprocessable Entity"})
		return
	}
	data.Frequency, err = strconv.Atoi(c.PostForm("frequency"))
	if err != nil {
		c.JSON(422, gin.H{"code": "422", "message": "Unprocessable Entity"})
		return
	}
	data.Failurethreshold, err = strconv.Atoi(c.PostForm("failure_threshold"))
	if err != nil {
		c.JSON(422, gin.H{"code": "422", "message": "Unprocessable Entity"})
		return
	}

	log.Printf("Checking for the health of requested URL...")
	if helper.Check(data.URL) {
		data.Status = "active"
	} else {
		data.Status = "inactive"
	}
	data.Activate = true

	log.Printf("Entering data into the database...")
	data.ID, err = db.InsertURL(&data)
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "message": fmt.Sprintf("%v", err)})
		return
	}
	varLock.Lock()
	UIDinfo[data.UUID] = INFO{data.Status, data.ID, data.Activate}
	varLock.Unlock()
	log.Printf("Initiating a frequency check...")
	go helper.PeriodicCheck(data.ID)
	log.Printf("Creating response object...")
	c.JSON(200, helper.CreateResponse(&data))
}
