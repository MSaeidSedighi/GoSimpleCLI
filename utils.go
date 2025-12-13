package main

import (
	"fmt"
	"time"

	"GO/notification"
)

type User struct {
	username string
	password string
	credit   int
}

func addUser(username, password string, notifier notification.Notification) {
	for _, user := range users {
		if user.username == username {
			fmt.Println("ERROR: this username exists!")
			return
		}
	}
	users = append(users, User{
		username: username,
		password: password,
		credit:   0,
	})
	// fmt.Println("New user added successfully.")
	notifier.Send("A new user is now added.")
}

func authenticate(username, password string, notifier notification.Notification) int {
	for id, user := range users {
		if user.username == username && user.password == password {
			// fmt.Println("Successful login.")
			notifier.Send(fmt.Sprintf("Login successful. Welcome %v", users[id].username))
			return id
		} else if user.username == username {
			fmt.Println("ERROR: Wrong password!")
			return -1
		}
	}
	fmt.Println("ERROR: User does not exist!")
	return -1
}

func withdraw(amount, userId int, notifier notification.Notification) {
	time.Sleep(20 * time.Second)
	if amount <= 0 {
		fmt.Println("ERROR: invalid amount (tip: amount > 0)")
		return
	}
	if users[userId].credit < amount {
		fmt.Printf("ERROR: Insufficient credit: %v.\n", users[userId].credit)
		return
	}
	users[userId].credit -= amount
	notifier.Send(fmt.Sprintf("A new withdraw! User %v's new credit is now %v.", users[userId].username, users[userId].credit))
	// fmt.Println("Successful withdraw")
}

func deposit(amount, userId int, notifier notification.Notification) {
	time.Sleep(20 * time.Second)
	if amount <= 0 {
		fmt.Println("ERROR: invalid amount (tip: amount > 0)")
		return
	}
	users[userId].credit += amount
	notifier.Send(fmt.Sprintf("A new deposit! User %v's new credit is now %v.", users[userId].username, users[userId].credit))
	// fmt.Println("Successful deposit.")
}
