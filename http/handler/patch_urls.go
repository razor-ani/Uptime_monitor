package handler

import (
	"Uptime-monitor/db"
	"Uptime-monitor/helper"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

//PatchUrls ...
func PatchUrls(c *gin.Context) {

	data := db.EntryData{}
	log.Printf("Accessing request params...")
	var err error
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
	data.UUID = c.Param("id")

	varLock.Lock()
	data.ID = UIDinfo[data.UUID].ID
	data.Status = UIDinfo[data.UUID].STATUS
	varLock.Unlock()
	if data.ID == 0 {
		c.JSON(400, gin.H{"code": "400", "message": "Invalid UUID"})
		return
	}
	log.Printf("Entering data into the database...")
	err = db.UpdateEntry(&data)
	if err != nil {
		c.JSON(500, gin.H{"code": "500", "message": fmt.Sprintf("%v", err)})
		return
	}
	log.Printf("Creating response object...")
	c.JSON(200, helper.CreateResponse(&data))
}
