package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	username string
	password string
	credit   int
}

var users = []User{}

var wg = sync.WaitGroup{}

func addUser(username, password string) {
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
	fmt.Println("New user added successfully.")
}

func authenticate(username, password string) int {
	for id, user := range users {
		if user.username == username && user.password == password {
			fmt.Println("Successful login.")
			return id
		} else if user.username == username {
			fmt.Println("ERROR: Wrong password!")
			return -1
		}
	}
	fmt.Println("ERROR: User does not exist!")
	return -1
}

func withdraw(amount, userId int) {
	if amount <= 0 {
		fmt.Println("ERROR: invalid amount (tip: amount > 0)")
		return
	}
	if users[userId].credit < amount {
		fmt.Printf("ERROR: Insufficient credit: %v.\n", users[userId].credit)
		return
	}
	users[userId].credit -= amount
	fmt.Println("Successful withdraw")
}

func deposit(amount, userId int) {
	if amount <= 0 {
		fmt.Println("ERROR: invalid amount (tip: amount > 0)")
		return
	}
	users[userId].credit += amount
	fmt.Println("Successful deposit.")
}

func main() {
	fmt.Println("Welcome to CLI")
	var isLoggedIn = false
	var id int
	reader := bufio.NewReader(os.Stdin)

	for {
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
				fmt.Println("ERROR: User are already logged in.")
				continue
			}
			addUser(fields[1], fields[2])
		case "login":
			// login
			if len(fields) < 3 {
				fmt.Println("ERROR: Missing argument! (tip: login <username> <password>)")
				continue
			}
			if isLoggedIn {
				fmt.Println("ERROR: User are already logged in.")
				continue
			}
			result := authenticate(fields[1], fields[2])
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
			deposit(amount, id)
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
			withdraw(amount, id)
		case "users":
			// see users
			fmt.Println(users)
		case "credit":
			// see credit
			if !isLoggedIn {
				fmt.Println("ERROR: Log in first!")
				continue
			}
			fmt.Printf("Your credit: %v.\n", users[id].credit)
		default:
			fmt.Println("ERROR: command not defined!")
		}
	}
}
