package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	content, _ := ioutil.ReadFile("./ProvisioningWSService.wsdl")
	w.Header().Add("Content-Type", "application/soap+xml")
	w.Write([]byte(content))
	fmt.Println("1")
}

func main() {
	http.HandleFunc("/wsdl", handleGet)
	http.ListenAndServe(":8000", nil)
}
