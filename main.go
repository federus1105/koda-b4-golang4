package main

import (
	"day2golang/internals"
	"fmt"
	"os"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error:", r)
			main()
		}
	}()

	a := &internals.AuthSystem{Users: make(map[string]*internals.User)}

	for {
		fmt.Println("\n====> Hi, Welcome to Sistem <===")
		fmt.Println("\n1.Register")
		fmt.Println("2.Login")
		fmt.Println("3.Forgot Password")
		fmt.Println("\n0.Exit")
		fmt.Print("Choose a Menu: ")

		var input string
		fmt.Scanln(&input)

		switch input {
		case "1":
			a.Register()

		case "2":
			a.Login()

		case "0":
			os.Exit(0)

		default:
			fmt.Println("Wrong selection!")
		}

		fmt.Println()
	}
}
