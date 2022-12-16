package main

import (
	"fmt"
)

func main() {
outer:
	for {
		fmt.Println("\n=====================")
		fmt.Println("Welcome to RideShare!\n",
			"1. Passenger Login\n",
			"2. Driver Login\n",
			"3. Create Account\n",
			"0. Quit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1: //passenger login
		case 2: //driver login
		case 3: //create account
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
			} else {
				createDriver()
			}
		case 0: //quit
			fmt.Println("Thank you for using RideShare, Goodbye!")
			break outer
		}
	}
}

func createPassenger() {
	fmt.Println("passenger success")
}

func createDriver() {
	fmt.Println("driver success")
}
