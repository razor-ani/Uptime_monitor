package helper

import (
	"Uptime-monitor/db"
)

//CreateResponse ...
func CreateResponse(d *db.EntryData) map[string]string {

	return map[string]string{
		"id":                string(d.ID),
		"url":               d.URL,
		"crawl_timeout":     string(d.Crawltimeout),
		"frequency":         string(d.Frequency),
		"failure_threshold": string(d.Failurethreshold),
		"status":            d.Status,
		"failure_count":     string(d.Failurecount),
	}
}
