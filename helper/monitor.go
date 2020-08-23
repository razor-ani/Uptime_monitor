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
func PeriodicCheck(id uint64) error {
	d, e := db.FetchAllData(id)
	if e != nil {
		return fmt.Errorf("Bad Gateway: %v", e)
	}
	if d.Activate {
		resp, err := http.Get("http://" + d.URL)
		if err != nil {
			d.Failurecount = d.Failurecount + 1
			if d.Failurecount == d.Failurethreshold {
				d.Status = "inactive"
				d.Failurecount = 0
			}
		}

		if string(resp.Status) == "200 OK" {
			d.Status = "active"
			d.Failurecount = 0
		}
		err = db.PeriodicUpdata(d)
		if err != nil {
			return fmt.Errorf("Bad Gateway: %v", err)
		}
		time.Sleep(time.Second * time.Duration(d.Frequency))
		return PeriodicCheck(d.ID)
	}
	return nil
}
