package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yahoo/vssh"
)

func main() {
	vs := vssh.New().Start()
	config := vssh.GetConfigUserPass("root", "zdns@knet.cn")
	for _, addr := range []string{"192.168.134.231:22", "192.168.134.232:22"} {
		vs.AddClient(addr, config, vssh.SetMaxSessions(4))
	}
	vs.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := "edig default www.qq.com"
	timeout, _ := time.ParseDuration("6s")
	respChan := vs.Run(ctx, cmd, timeout)

	for resp := range respChan {
		if err := resp.Err(); err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(resp.ID())
		outTxt, errTxt, _ := resp.GetText(vs)
		fmt.Println(outTxt, errTxt, resp.ExitStatus())
	}
}
