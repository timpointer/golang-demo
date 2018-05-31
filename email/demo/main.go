package main

import (
	"log"

	gomail "gopkg.in/gomail.v2"
)

func main() {

	m := gomail.NewMessage()

	m.SetHeader("From", "tjing@evolveconsulting.com.hk")
	m.SetHeader("To", "adong@evolveconsulting.com.hk")
	m.SetHeader("Subject", "3v报表")
	m.SetBody("text/html", "请查收报表!")
	m.Attach("test.txt")

	d := gomail.NewDialer("smtp.office365.com", 587, "tjing@evolveconsulting.com.hk", "password")
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Printf("dialAndSend :%v", err)
	}
}
