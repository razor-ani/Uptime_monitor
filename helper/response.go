package helper

import (
	"Uptime-monitor/db"
	"strconv"
)

//CreateResponse ...
func CreateResponse(d *db.EntryData) map[string]string {

	return map[string]string{
		"id":                strconv.Itoa(int(d.ID)),
		"uuid":              d.UUID,
		"url":               d.URL,
		"crawl_timeout":     strconv.Itoa(d.Crawltimeout),
		"frequency":         strconv.Itoa(d.Frequency),
		"failure_threshold": strconv.Itoa(d.Failurethreshold),
		"status":            d.Status,
		"failure_count":     strconv.Itoa(d.Failurecount),
	}
}
