package db

import (
	"database/sql"
	"errors"
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
func Getdb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "uptime:UPmonitor11@tcp(127.0.0.1:3306)/uptime_db")

	if err != nil {
		return nil, err
	}
	return db, nil

}

//InsertURL  ..
func InsertURL(b *EntryData) (uint64, error) {

	dbLock.Lock()
	log.Printf("Fetching DB instance...")
	d, e := Getdb()
	if e != nil {
		return 0, e
	}

	stmt, err := d.Prepare("INSERT INTO url_data ( uuid, url, frequency, crawl_timeout, failure_threshold, status,failure_count, activate) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return 0, err
	}
	res, err := stmt.Exec(b.UUID, b.URL, b.Frequency, b.Crawltimeout, b.Failurethreshold, b.Status, b.Failurecount, b.Activate)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer d.Close()
	defer dbLock.Unlock()
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	b.ID = uint64(lastID)

	return b.ID, nil

}

//FetchURLInfo ..
func FetchURLInfo(id uint64) (string, error) {
	dbLock.Lock()
	log.Printf("Fetching DB instance...")
	d, e := Getdb()
	if e != nil {
		return "", e
	}
	resp, err := d.Query("SELECT status FROM url_data WHERE id=" + strconv.Itoa(int(id)))
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer d.Close()
	defer dbLock.Unlock()

	var status string
	if resp.Next() {
		resp.Scan(&status)
	} else {
		return "", errors.New("Couldn't found data")
	}

	return status, nil

}

//UpdateEntry ...
func UpdateEntry(b *EntryData) error {
	dbLock.Lock()
	log.Printf("Fetching DB instance...")
	d, e := Getdb()
	if e != nil {
		return e
	}
	stmt, err := d.Prepare("UPDATE url_data SET frequency = ? , failure_threshold = ?, crawl_timeout = ? WHERE id=?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(b.Frequency, b.Failurethreshold, b.Crawltimeout, b.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Patch Successful!")

	resp, err := d.Query("SELECT url, activate FROM url_data WHERE id=" + strconv.Itoa(int(b.ID)))
	if err != nil {
		log.Println(err)
		return err
	}

	defer d.Close()
	defer dbLock.Unlock()

	if resp.Next() {
		resp.Scan(&b.URL, &b.Activate)
	} else {
		return errors.New("Couldn't found data")
	}
	return nil
}

//FetchAllData ..
func FetchAllData(id uint64) (*EntryData, error) {

	dbLock.Lock()

	d, e := Getdb()
	if e != nil {
		return nil, e
	}

	data := EntryData{}

	resp, err := d.Query("SELECT id, uuid, url, crawl_timeout, frequency, failure_threshold, status, failure_count, activate FROM url_data WHERE id=" + strconv.Itoa(int(id)))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer d.Close()
	defer dbLock.Unlock()

	if resp.Next() {
		resp.Scan(&data.ID, &data.UUID, &data.URL, &data.Crawltimeout, &data.Frequency, &data.Failurethreshold, &data.Status, &data.Failurecount, &data.Activate)
	} else {
		return nil, errors.New("Couldn't found data")
	}

	return &data, nil

}

//PeriodicUpdata ..
func PeriodicUpdata(b *EntryData) error {
	dbLock.Lock()

	d, e := Getdb()
	if e != nil {
		return e
	}

	stmt, err := d.Prepare("UPDATE url_data SET failure_count = ? , status = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(b.Failurecount, b.Status, b.ID)
	if err != nil {
		return err
	}
	defer d.Close()
	defer dbLock.Unlock()

	if err != nil {
		return err
	}
	return nil

}

//UpdateActivation ..
func UpdateActivation(id uint64, a bool) error {

	dbLock.Lock()
	log.Printf("Fetching DB instance...")
	d, e := Getdb()
	if e != nil {
		return e
	}
	log.Printf("Fetching Data...")
	stmt, err := d.Prepare("UPDATE url_data SET activate = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(a, id)

	if err != nil {
		return err
	}

	defer d.Close()
	defer dbLock.Unlock()
	if err != nil {
		return err
	}
	return nil
}

//DeleteWithID ..
func DeleteWithID(id string) error {

	dbLock.Lock()
	log.Printf("Fetching DB instance...")
	d, e := Getdb()
	if e != nil {
		return e
	}

	_, err := d.Query("DELETE FROM url_data  WHERE id =" + id)
	if err != nil {
		return err
	}
	defer d.Close()
	defer dbLock.Unlock()

	return nil
}
