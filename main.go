package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"GO/notification"
)

var users = []User{}

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Welcome to CLI")
	var isLoggedIn = false
	var id int
	var isExitCommand = false
	reader := bufio.NewReader(os.Stdin)
	sms_notifier := notification.SMSNotif{}
	email_notifier := notification.EmailNotif{}

	for {
		if isExitCommand {
			break
		}
		if isLoggedIn {
			fmt.Printf("%v->", users[id].username)
		} else {
			fmt.Print("->")
		}
		entry, _ := reader.ReadString('\n')
		entry = strings.TrimSpace(entry)

		fields := strings.Fields(entry)
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "signup":
			// signup
			if len(fields) < 3 {
				fmt.Println("ERROR: Missing argument! (tip: signup <username> <password>)")
				continue
			}
			if isLoggedIn {
				fmt.Println("ERROR: User is already logged in.")
				continue
			}
			addUser(fields[1], fields[2], sms_notifier)
		case "login":
			// login
			if len(fields) < 3 {
				fmt.Println("ERROR: Missing argument! (tip: login <username> <password>)")
				continue
			}
			if isLoggedIn {
				fmt.Println("ERROR: User is already logged in.")
				continue
			}
			result := authenticate(fields[1], fields[2], sms_notifier)
			if result != -1 {
				isLoggedIn = true
				id = result
			}
		case "logout":
			// logout
			isLoggedIn = false
			fmt.Println("Logged out")
		case "deposit":
			// deposit
			if !isLoggedIn {
				fmt.Println("ERROR: Log in first!")
				continue
			}
			if len(fields) < 2 {
				fmt.Println("ERROR: Missing argument! (tip: deposit <amount>)")
				continue
			}
			amount, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println(err)
			}
			wg.Add(1)
			go deposit(amount, id, email_notifier)
		case "withdraw":
			// withdraw
			if !isLoggedIn {
				fmt.Println("ERROR: Log in first!")
				continue
			}
			if len(fields) < 2 {
				fmt.Println("ERROR: Missing argument! (tip: withdraw <amount>)")
				continue
			}
			amount, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			wg.Add(1)
			go withdraw(amount, id, email_notifier)
		case "users":
			// see users
			fmt.Printf("%+v\n", users)
		case "credit":
			// see credit
			if !isLoggedIn {
				fmt.Println("ERROR: Log in first!")
				continue
			}
			fmt.Printf("Your credit: %v.\n", users[id].credit)
		case "exit":
			isExitCommand = true
		default:
			fmt.Println("ERROR: command not defined!")
		}
	}
	wg.Wait()
}
