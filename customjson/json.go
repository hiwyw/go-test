package main

import (
	"encoding/json"
	"fmt"
)

const (
	zdnsTrueStr  = `"yes"`
	zdnsFalseStr = `"no"`
)

type ZDNSBool bool

func (b ZDNSBool) MarshalJSON() ([]byte, error) {
	if b == true {
		return []byte(zdnsTrueStr), nil
	}
	return []byte(zdnsFalseStr), nil
}

func (b *ZDNSBool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case zdnsTrueStr:
		*b = true
		return nil
	case zdnsFalseStr:
		*b = false
		return nil
	default:
		*b = false
		return fmt.Errorf("ZDNSBool unknown value: %v", string(data))
	}
}

type RR struct {
	Domain   string   `json:"domain"`
	IsEnable ZDNSBool `json:"is_enable"`
}

func main() {
	r := RR{"www.test.com.", true}

	out, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	r2 := RR{}
	t := `{"domain": "www.test.com.", "is_enable": "no1"}`
	if err := json.Unmarshal([]byte(t), &r2); err != nil {
		panic(err)
	}
	fmt.Println(r2)
}
