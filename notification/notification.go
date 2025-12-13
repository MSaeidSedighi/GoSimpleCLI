package notification

import "fmt"

type Notification interface {
	Send(msg string)
}

type SMSNotif struct {
}

func (SMSNotif) Send(msg string) {
	fmt.Println("SMS:", msg)
}

type EmailNotif struct {
}

func (EmailNotif) Send(msg string) {
	fmt.Println("Email:", msg)
}
