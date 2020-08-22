package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"

	//mysql ...
	_ "github.com/go-sql-driver/mysql"
)

//dbLock ..
var dbLock sync.Mutex

//EntryData is ... id, uuid, url, crawl_timeout, frequency, failure_threshold, status, failure_count, activate
type EntryData struct {
	ID               uint64 `json:"id"`
	UUID             string `json:"uuid"`
	URL              string `json:"url"`
	Crawltimeout     int    `json:"crawl_timeout"`
	Frequency        int    `json:"frequency"`
	Failurethreshold int    `json:"failure_threshold"`
	Status           string `json:"status"`
	Failurecount     int    `json:"failure_count"`
	Activate         bool   `json:"activate"`
}

const dbUser = "uptime"
const pass = "UPmonitor11"
const dbName = "uptime_db"
const tableName = "url_data"

//Getdb is ..
func Getdb() *sql.DB {
	db, err := sql.Open("mysql", "uptime:UPmonitor11@tcp(127.0.0.1:3306)/uptime_db")

	if err != nil {
		panic(err)
	} else {
		return db
	}

}

//InsertURL  ..
func InsertURL(b *EntryData) uint64 {
	dbLock.Lock()
	d := Getdb()
	stmt, err := d.Prepare("INSERT INTO url_data ( uuid, url, frequency, crawl_timeout, failure_threshold, status,failure_count, activate) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(b.UUID, b.URL, b.Frequency, b.Crawltimeout, b.Failurethreshold, b.Status, b.Failurecount, b.Activate)
	defer d.Close()
	defer dbLock.Unlock()
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	b.ID = uint64(lastID)

	return b.ID

}

//FetchURLInfo ..
func FetchURLInfo(id uint64) string {
	dbLock.Lock()
	d := Getdb()

	resp, err := d.Query("SELECT status FROM url_data WHERE id=" + strconv.Itoa(int(id)))
	defer d.Close()
	defer dbLock.Unlock()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if resp.Next() {
		resp.Scan(&status)
	}

	return status

}

//UpdateEntry ...
func UpdateEntry(b *EntryData) {
	dbLock.Lock()
	d := Getdb()

	stmt, err := d.Prepare("UPDATE url_data SET frequency = ? , failure_threshold = ?, crawl_timeout = ? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(b.Frequency, b.Failurethreshold, b.Crawltimeout, b.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Patch Successful!")

	resp, err := d.Query("SELECT url, activate FROM url_data WHERE id=" + strconv.Itoa(int(b.ID)))

	defer d.Close()
	defer dbLock.Unlock()

	if err != nil {
		log.Fatal(err)
	}
	if resp.Next() {
		resp.Scan(&b.URL, &b.Activate)
	}

}

//FetchAllData ..
func FetchAllData(id uint64) *EntryData {

	dbLock.Lock()
	d := Getdb()
	data := EntryData{}

	resp, err := d.Query("SELECT id, uuid, url, crawl_timeout, frequency, failure_threshold, status, failure_count, activate FROM url_data WHERE id=" + strconv.Itoa(int(id)))
	defer d.Close()
	defer dbLock.Unlock()

	if err != nil {
		log.Fatal(err)
	}

	if resp.Next() {
		resp.Scan(&data.ID, &data.UUID, &data.URL, &data.Crawltimeout, &data.Frequency, &data.Failurethreshold, &data.Status, &data.Failurecount, &data.Activate)
	}

	return &data

}

//PeriodicUpdata ..
func PeriodicUpdata(b *EntryData) {
	dbLock.Lock()
	d := Getdb()

	stmt, err := d.Prepare("UPDATE url_data SET failure_count = ? , status = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(b.Failurecount, b.Status, b.ID)
	defer d.Close()
	defer dbLock.Unlock()

	if err != nil {
		log.Fatal(err)
	}

}

//UpdateActivation ..
func UpdateActivation(id uint64, a bool) {

	dbLock.Lock()
	d := Getdb()

	stmt, err := d.Prepare("UPDATE url_data SET activate = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(a, id)

	defer d.Close()
	defer dbLock.Unlock()
	if err != nil {
		log.Fatal(err)
	}
}

//DeleteWithID ..
func DeleteWithID(id string) {

	dbLock.Lock()
	d := Getdb()

	_, err := d.Query("DELETE FROM url_data  WHERE id =" + id)
	defer d.Close()
	defer dbLock.Unlock()

	if err != nil {
		log.Fatal(err)
	}
}
