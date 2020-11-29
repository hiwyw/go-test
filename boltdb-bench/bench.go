package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/boltdb/bolt"
)

const (
	domainZone         = "test.com"
	domainResponseIP   = "1.1.1.1"
	domainServerRoomID = "300100000001"

	dbPath = "domain.db"
)

var db *bolt.DB

func main() {
	// begin := time.Now()
	initDB()
	// addDomainDataToDB()
	// log.Printf("add 1000000 key to db cost %s\n", time.Since(begin))
	// for i := 0; i < 100000; i++ {
	// 	err := addOneDomainToDB(genDomainKey(i))
	// 	if err != nil {
	// 		log.Printf("add %s failed %s", genDomainKey(i))
	// 	}
	// }
	// log.Printf("add 100000 key cost %s", time.Since(begin))
	for i := 0; i < 100; i++ {
		printDomainKey(genDomainKey(i))
	}
}

func printDomainKey(key string) {
	db.View(func(t *bolt.Tx) error {
		b := t.Bucket([]byte("domains"))
		v := b.Get([]byte(key))
		log.Printf("%s==>%s", key, string(v))
		return nil
	})
}

func addOneDomainToDB(key string) error {
	return db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte("domains"))
		v := b.Get([]byte(key))
		vInt, _ := strconv.Atoi(string(v))
		vInt++
		b.Put([]byte(key), []byte(strconv.Itoa(vInt)))
		return nil
	})
}

func genDomainKey(index int) string {
	return fmt.Sprintf("%d.%s|%s|%s", index, domainZone, domainResponseIP, domainServerRoomID)
}

func initDB() {
	d, err := bolt.Open(dbPath, 0644, nil)
	if err != nil {
		log.Fatalf("init db failed %s\n", err)
	}
	db = d
	err = db.Update(func(t *bolt.Tx) error {
		_, err := t.CreateBucketIfNotExists([]byte("domains"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalf("init db failed %s\n", err)
	}
}

func addDomainDataToDB() {
	err := db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte("domains"))
		for i := 0; i < 500000; i++ {
			k := genDomainKey(i)
			err := b.Put([]byte(k), []byte("1"))
			if err != nil {
				log.Printf("put %s failed %s", k, err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("add domain data to db failed %s", err)
	}
}
