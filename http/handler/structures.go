package handler

import (
	"sync"
)

//PostURLReq ..
type PostURLReq struct {
	URL              string `json:"url"`
	Crawltimeout     int    `json:"crawl_timeout"`
	Frequency        int    `json:"frequency"`
	Failurethreshold int    `json:"frequency_threshold"`
}

//varLock ..
var varLock sync.Mutex

//INFO ..
type INFO struct {
	STATUS     string
	ID         uint64
	ACTIVATION bool
}

//UIDinfo ..
var UIDinfo map[string]INFO = make(map[string]INFO)
