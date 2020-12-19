package main

import "gopkg.in/gomail.v2"

func main() {
	if err := sendMail(); err != nil {
		panic(err)
	}
}

func sendMail() error {
	m := gomail.NewMessage()
	m.SetHeader("From", "wyw_314@126.com")
	m.SetHeader("To", "17319330797@189.cn")
	m.SetHeader("Subject", "gomail")
	m.SetBody("text/html", "this is a test mail")

	d := gomail.NewDialer("smtp.126.com", 25, "wyw_314@126.com", "UEJSBOXOLVOTHVSU")
	return d.DialAndSend(m)
}
