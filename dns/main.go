package main

import (
	"fmt"

	"github.com/miekg/dns"
)

func main() {
	cli := new(dns.Client)

	msg := new(dns.Msg)
	msg.SetQuestion("www.qq.com.", dns.TypeA)
	msg.RecursionDesired = true

	r, rtt, err := cli.Exchange(msg, "114.114.114.114:53")
	if err != nil {
		panic(err)
	}

	for _, rr := range r.Answer {
		fmt.Println(rr.Header().Ttl)
	}

	fmt.Printf("rtt==>%v\n", rtt)
	fmt.Printf("response==>%v\n", r)
}
