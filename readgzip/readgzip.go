package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const gzipFile = "query.log..0.gz"

func main() {
	begin := time.Now()

	f, err := os.Open(gzipFile)
	if err != nil {
		log.Fatalf("open gzip file failed %s", err.Error())
		return
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	buf := bytes.NewBuffer(data)

	zr, err := gzip.NewReader(buf)
	if err != nil {
		log.Fatalf("new gzip reader failed %s", err.Error())
	}
	defer zr.Close()

	scanner := bufio.NewScanner(zr)
	for {
		ok := scanner.Scan()
		if ok == false {
			err := scanner.Err()
			if err == nil {
				log.Printf("reached EOF")
				break
			}
			log.Fatal(err)
		}
		log.Print(scanner.Text())
	}

	cost := time.Since(begin)
	log.Println("cost time: ", cost.Microseconds())
}
