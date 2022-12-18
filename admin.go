package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Passenger struct {
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	PhoneNo   int    `json:"PhoneNo"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
}

type Driver struct {
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	PhoneNo   int    `json:"PhoneNo"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	IDNo      string `json:"IDNo"`
	LicenseNo string `json:"LicenseNo"`
	IsBusy    string `json:"IsBusy"`
}

type CarTrip struct {
	ID                int    `json:"ID"`
	PassengerUsername string `json:"PassengerUsername"`
	DriverUsername    string `json:"DriverUsername"`
	Pickup            string `json:"Pickup"`
	Dropoff           string `json:"Dropoff"`
	PickupTime        string `json:"PickupTime"`
	IsCompleted       string `json:"IsCompleted"`
}

type CarTrips struct {
	CarTrips map[string]CarTrip `json:"CarTrips"`
}

func main() {
outer:
	for {
		fmt.Println("\n=====================")
		fmt.Println("Welcome to RideShare!\n",
			"1. Create Account\n",
			"2. Update Passenger Information\n",
			"3. Update Driver Information\n",
			"4. Request trip\n",
			"6. Display all trips taken\n",
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
				"0. Back")
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
			updateDriver()
		case 4: //add car trip
			createCarTrip()
		case 6: //display all trips taken by a passenger
			displayAllTrips()
		case 0: //quit
			fmt.Println("Thank you for using RideShare, Goodbye!")
			break outer
		}
	}
}

func displayAllTrips() {
	var newCarTrip CarTrip

	fmt.Print("Enter Username: ") //retrieve username to know which passenger car trips are to be displayed
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newCarTrip.PassengerUsername = strings.TrimSpace(input)

	jsonString, _ := json.Marshal(newCarTrip)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5000/api/v1/cartrip/"+newCarTrip.PassengerUsername, resbody); err == nil {
		if res, err := client.Do(req); err == nil {
			if body, err := ioutil.ReadAll(res.Body); err == nil {
				if res.StatusCode == 404 { //checks if there is an error caused by invalid username
					fmt.Println("Error - Username does not exists!")
				} else {
					var res CarTrips
					json.Unmarshal(body, &res)

					fmt.Print("\n=====================")
					for _, v := range res.CarTrips {
						fmt.Println("\nDriver's Username:", v.DriverUsername)
						fmt.Println("Pickup:", v.Pickup)
						fmt.Println("Dropoff:", v.Dropoff)
						fmt.Println("Pickup Time:", v.PickupTime)
					}
				}
			}
		}
	}
}

func createCarTrip() { //create a new Car Trip record
	var newCarTrip CarTrip

	fmt.Print("Enter Username: ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	newCarTrip.PassengerUsername = strings.TrimSpace(input0)

	newCarTrip.DriverUsername = ""

	fmt.Print("Enter Postal Code for Pick-up: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newCarTrip.Pickup = strings.TrimSpace(input)

	fmt.Print("Enter Postal Code for Drop-off: ")
	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	newCarTrip.Dropoff = strings.TrimSpace(input2)

	newCarTrip.PickupTime = time.Now().Format("2006-01-02 15:04:05")
	newCarTrip.IsCompleted = "0"

	jsonString, _ := json.Marshal(newCarTrip)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/cartrip/"+newCarTrip.PassengerUsername, resbody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Request for ride succesfully created by", newCarTrip.PassengerUsername)
			} else {
				fmt.Println("Error - Username does not exists!")
			}
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
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

func createDriver() { //create a new Driver Account
	var newDriver Driver

	fmt.Print("Enter Username: ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	newDriver.Username = strings.TrimSpace(input0)

	fmt.Print("Enter First Name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newDriver.FirstName = strings.TrimSpace(input)

	fmt.Print("Enter Last Name: ")
	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	newDriver.LastName = strings.TrimSpace(input2)

	fmt.Print("Enter Phone Number: ")
	fmt.Scanf("%d\n", &(newDriver.PhoneNo))

	fmt.Print("Enter Email: ")
	reader3 := bufio.NewReader(os.Stdin)
	input3, _ := reader3.ReadString('\n')
	newDriver.Email = strings.TrimSpace(input3)

	fmt.Print("Enter Password: ")
	reader4 := bufio.NewReader(os.Stdin)
	input4, _ := reader4.ReadString('\n')
	newDriver.Password = strings.TrimSpace(input4)

	fmt.Print("Enter NRIC: ")
	reader5 := bufio.NewReader(os.Stdin)
	input5, _ := reader5.ReadString('\n')
	newDriver.IDNo = strings.TrimSpace(input5)

	fmt.Print("Enter Car License Number: ")
	reader6 := bufio.NewReader(os.Stdin)
	input6, _ := reader6.ReadString('\n')
	newDriver.LicenseNo = strings.TrimSpace(input6)

	newDriver.IsBusy = "0"

	jsonString, _ := json.Marshal(newDriver)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5000/api/v1/driver/"+newDriver.Username, resbody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Account for", newDriver.Username, "created successfully")
			} else {
				fmt.Println("Error - Username exists!")
			}
		}
	}
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

func updateDriver() { //update an existing Driver Account
	var newDriver Driver

	fmt.Print("Enter Username: ")
	reader0 := bufio.NewReader(os.Stdin)
	input0, _ := reader0.ReadString('\n')
	newDriver.Username = strings.TrimSpace(input0)

	fmt.Print("Update First Name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newDriver.FirstName = strings.TrimSpace(input)

	fmt.Print("Update Last Name: ")
	reader2 := bufio.NewReader(os.Stdin)
	input2, _ := reader2.ReadString('\n')
	newDriver.LastName = strings.TrimSpace(input2)

	fmt.Print("Update Phone Number: ")
	fmt.Scanf("%d\n", &(newDriver.PhoneNo))

	fmt.Print("Update Email: ")
	reader3 := bufio.NewReader(os.Stdin)
	input3, _ := reader3.ReadString('\n')
	newDriver.Email = strings.TrimSpace(input3)

	fmt.Print("Update Password: ")
	reader4 := bufio.NewReader(os.Stdin)
	input4, _ := reader4.ReadString('\n')
	newDriver.Password = strings.TrimSpace(input4)

	fmt.Print("Update Car License Number: ")
	reader6 := bufio.NewReader(os.Stdin)
	input6, _ := reader6.ReadString('\n')
	newDriver.LicenseNo = strings.TrimSpace(input6)

	jsonString, _ := json.Marshal(newDriver)
	resbody := bytes.NewBuffer(jsonString)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:5000/api/v1/driver/"+newDriver.Username, resbody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Account for", newDriver.Username, "updated successfully")
			} else {
				fmt.Println("Error - Username does not exists!")
			}
		}
	}
}
