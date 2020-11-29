package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	domainZone         = "test.com"
	domainResponseIP   = "1.1.1.1"
	domainServerRoomID = "300100000001"

	dbPath = "domain.db"
)

func main() {
	begin := time.Now()
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("init db failed %s", err)
	}

	// for i := 0; i < 5000000; i++ {
	// 	db.Put([]byte(genDomainKey(i)), []byte("1"), nil)
	// }
	for i := 0; i < 2; i++ {
		for j := 0; j < 5000000; j++ {
			k := genDomainKey(j)
			v, err := db.Get([]byte(k), nil)
			if err != nil {
				log.Printf("get %s failed %s\n", k, err)
			}
			vInt, _ := strconv.Atoi(string(v))
			vInt++
			db.Put([]byte(k), []byte(strconv.Itoa(vInt)), nil)
		}
	}

	for i := 0; i < 100; i++ {
		k := genDomainKey(i)
		v, _ := db.Get([]byte(k), nil)
		log.Printf("%s==>%s", k, string(v))
	}
	log.Printf("add 500 0000 key cost %s\n", time.Since(begin))
}

func genDomainKey(index int) string {
	return fmt.Sprintf("%d.%s|%s|%s", index, domainZone, domainResponseIP, domainServerRoomID)
}
