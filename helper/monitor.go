package helper

import (
	"Uptime-monitor/db"
	"fmt"
	"net/http"
	"time"
)

//Check ..
func Check(url string) bool {
	resp, error := http.Get("http://" + url)
	if error != nil {
		fmt.Printf("External server errror: %v \n", error)
		return false
	}
	return string(resp.Status) == "200 OK"
}

//PeriodicCheck ...
func PeriodicCheck(id uint64) {
	d := db.FetchAllData(id)
	if d.Activate {
		resp, error := http.Get("http://" + d.URL)
		if error != nil {
			d.Failurecount = d.Failurecount + 1
			if d.Failurecount == d.Failurethreshold {
				d.Status = "inactive"
				d.Failurecount = 0
			}
		}

		if string(resp.Status) == "200 OK" {
			d.Status = "active"
		}
		db.PeriodicUpdata(d)
		time.Sleep(time.Second * time.Duration(d.Frequency))
		PeriodicCheck(d.ID)
	}
	return
}
