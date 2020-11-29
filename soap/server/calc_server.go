package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/hooklift/gowsdl/soap"

	"github.com/hiwyw/go-test/soap/calc"
)

func main() {
	s := NewSOAPServer("localhost:8080")
	log.Fatal(s.ListenAndServe())
}

func NewSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", soapHandler)
	return mux
}

func NewSOAPServer(addr string) *http.Server {
	mux := NewSOAPMux()
	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}
	return server
}

func soapHandler(w http.ResponseWriter, r *http.Request) {
	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// match method
	var res interface{}
	m := regexp.MustCompile(`<Add xmlns=`)
	if m.MatchString(string(rawBody)) {
		res = processAdd(rawBody)
	} else {
		res = nil
		fmt.Println("the method requested is not available")
	}

	v := soap.SOAPEnvelope{
		Body: soap.SOAPBody{
			Content: res,
		},
	}
	w.Header().Set("Content-Type", "text/xml")
	x, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(x)
	return
}

func processAdd(body []byte) *calc.AddResponse {
	envlop := &soap.SOAPEnvelope{
		Body: soap.SOAPBody{
			Content: &calc.Add{},
		},
	}
	err := xml.Unmarshal(body, envlop)
	if err != nil {
		fmt.Println("xml Unmarshal error:", err)
		return nil
	}

	fmt.Println(envlop.Body.Content)

	r, ok := envlop.Body.Content.(*calc.Add)
	if !ok {
		return nil
	} else {
		return &calc.AddResponse{
			AddResult: r.IntA + r.IntB,
		}
	}
}
