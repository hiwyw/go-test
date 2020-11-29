package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type RR struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	TTL   int    `json:"ttl"`
	Type  string `json:"type"`
	Rdata string `json:"rdata"`
}

func (r *RR) ToJson() string {
	out, _ := json.Marshal(r)
	return string(out)
}

var db *sql.DB

func main() {
	initDB()
	for i := 0; i < 10; i++ {
		r := &RR{
			Name:  fmt.Sprintf("%d.test.com.", i),
			TTL:   60,
			Type:  "A",
			Rdata: "1.1.1.1",
		}
		if err := add(r); err != nil {
			log.Fatalf("add rr failed %s", err.Error())
		}
	}

	rrs, err := list()
	if err != nil {
		log.Fatalf("list rr failed %s", err.Error())
	}
	out, _ := json.Marshal(rrs)
	fmt.Println(string(out))

	if err := delete(); err != nil {
		log.Fatalf("delete table failed %s", err.Error())
	}
}

func initDB() {
	d, err := sql.Open("sqlite3", "./rr.db")
	if err != nil {
		log.Fatalf("init db failed %s", err.Error())
	}
	db = d
}

func add(r *RR) error {
	sql := `INSERT INTO RR (name,ttl,type,rdata) VALUES (:1,:2,:3,:4);`
	_, err := db.Exec(sql, r.Name, r.TTL, r.Type, r.Rdata)
	return err
}

func list() ([]*RR, error) {
	sql := `SELECT * FROM RR;`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rrs := []*RR{}
	for rows.Next() {
		rr := &RR{}
		if err := rows.Scan(&rr.ID, &rr.Name, &rr.TTL, &rr.Type, &rr.Rdata); err != nil {
			return nil, err
		}
		rrs = append(rrs, rr)
	}

	return rrs, nil
}

func delete() error {
	sql := `DELETE FROM RR;`
	_, err := db.Exec(sql)
	return err
}
