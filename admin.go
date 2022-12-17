package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Passenger struct {
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	PhoneNo   int    `json:"PhoneNo"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}

func main() {
outer:
	for {
		fmt.Println("\n=====================")
		fmt.Println("Welcome to RideShare!\n",
			"1. Create Account\n",
			"2. Update Passenger Information\n",
			"3. Update Driver Information\n",
			"0. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //create account
			fmt.Println("\n=====================")
			fmt.Println("Choose User Type\n",
				"1. Create Passenger Account\n",
				"2. Create Driver Account\n",
				"0. Quit")
			fmt.Print("Enter an option: ")

			var choice2 int
			fmt.Scanf("%d\n", &choice2)

			if choice2 == 1 {
				createPassenger()
			} else if choice2 == 2 {
				createDriver()
			} else {
				break
			}
		case 2: //update passenger info
			updatePassenger()
		case 3: //update driver info
		case 0: //quit
			fmt.Println("Thank you for using RideShare, Goodbye!")
			break outer
		}
	}
}

func createPassenger() { //create a new Passenger Account
	var newPassenger Passenger

	fmt.Print("Enter Username: ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	newPassenger.Username = strings.TrimSpace(input0)

	fmt.Print("Enter First Name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newPassenger.FirstName = strings.TrimSpace(input)

	fmt.Print("Enter Last Name: ")
	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	newPassenger.LastName = strings.TrimSpace(input2)

	fmt.Print("Enter Phone Number: ")
	fmt.Scanf("%d\n", &(newPassenger.PhoneNo))

	fmt.Print("Enter Email: ")
	reader3 := bufio.NewReader(os.Stdin)
	input3, _ := reader3.ReadString('\n')
	newPassenger.Email = strings.TrimSpace(input3)

	fmt.Print("Enter Password: ")
	reader4 := bufio.NewReader(os.Stdin)
	input4, _ := reader4.ReadString('\n')
	newPassenger.Password = strings.TrimSpace(input4)

	jsonString, _ := json.Marshal(newPassenger)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/passenger/"+newPassenger.Username, resbody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Account for", newPassenger.Username, "created successfully")
			} else {
				fmt.Println("Error - Username exists!")
			}
		}
	}
}

func createDriver() {
	fmt.Println("driver success")
}

func updatePassenger() { //update an existing Passenger Account
	var newPassenger Passenger

	fmt.Print("Enter Username: ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	newPassenger.Username = strings.TrimSpace(input0)

	fmt.Print("Update First Name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newPassenger.FirstName = strings.TrimSpace(input)

	fmt.Print("Update Last Name: ")
	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	newPassenger.LastName = strings.TrimSpace(input2)

	fmt.Print("Update Phone Number: ")
	fmt.Scanf("%d\n", &(newPassenger.PhoneNo))

	fmt.Print("Update Email: ")
	reader3 := bufio.NewReader(os.Stdin)
	input3, _ := reader3.ReadString('\n')
	newPassenger.Email = strings.TrimSpace(input3)

	fmt.Print("Update Password: ")
	reader4 := bufio.NewReader(os.Stdin)
	input4, _ := reader4.ReadString('\n')
	newPassenger.Password = strings.TrimSpace(input4)

	jsonString, _ := json.Marshal(newPassenger)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/passenger/"+newPassenger.Username, resbody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Account for", newPassenger.Username, "updated successfully")
			} else {
				fmt.Println("Error - Username does not exists!")
			}
		}
	}
}
