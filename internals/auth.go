package internals

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AuthInterface interface {
	Register()
	Login()
	ForgotPassword()
}

type AuthSystem struct {
	Users map[string]*User
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (a *AuthSystem) Register() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("--- Register ---")

		fmt.Print("What is your first name: ")
		inputFirstname := readInput(reader)

		fmt.Print("What is your last name: ")
		inputLastname := readInput(reader)

		fmt.Print("What is your email: ")
		inputEmail := readInput(reader)
		if _, exists := a.Users[inputEmail]; exists {
			panic("Email is used!")
		}

		fmt.Print("Enter a strong password: ")
		inputPassword := readInput(reader)
		md5Hash := GenerateMD5Hash(inputPassword)
		_ = md5Hash

		fmt.Print("Confirm your password: ")
		inputConfirm := readInput(reader)

		if inputPassword != inputConfirm {
			fmt.Print("Wrong confirm password. Press Enter to try again.")
			readInput(reader)
			continue
		}

		fmt.Println("\nIs this information correct?")
		fmt.Println("Firstname :", inputFirstname)
		fmt.Println("Lastname  :", inputLastname)
		fmt.Println("Email     :", inputEmail)

		fmt.Print("Continue? (Y/n): ")
		confirm := strings.ToUpper(readInput(reader))

		if confirm == "Y" {
			a.Users[inputEmail] = &User{
				Firstname: inputFirstname,
				Lastname:  inputLastname,
				Email:     inputEmail,
				Password:  inputPassword,
			}
			fmt.Print("Register success, press enter to back..")
			readInput(reader)
			break
		}
	}
}

func (a *AuthSystem) Login() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("--- LOGIN ---")
		fmt.Print("Enter your Email: ")
		inputEmail := readInput(reader)

		fmt.Print("Enter your Password: ")
		inputPassword := readInput(reader)

		user, exists := a.Users[inputEmail]
		if !exists || user.Password != inputPassword {
			fmt.Print("Wrong email or password. Press enter to restart.")
			readInput(reader)
			continue
		}

		fmt.Print("Login success! Press enter to back..")
		readInput(reader)

		a.Dashboard(user)
		break
	}
}

func (a *AuthSystem) Dashboard(user *User) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("\nHello %s!\n", user.Firstname)
		fmt.Println("1. List All Users")
		fmt.Println("2. Logout")
		fmt.Println("\n0. Exit")
		fmt.Print("Choose a menu: ")

		choice := readInput(reader)

		switch choice {
		case "1":
			fmt.Println("\n--- List all users ---")
			for _, u := range a.Users {
				fmt.Printf("- firstname: %s lastname: %s email: (%s)\n", u.Firstname, u.Lastname, u.Email)
			}
		case "2":
			fmt.Print("Logout success, press enter to back..")
			readInput(reader)
			return
		case "0":
			os.Exit(0)
		default:
			fmt.Println("Invalid choice, press enter to back..")
			readInput(reader)
		}
	}
}

func (a *AuthSystem) ForgotPassword() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("--- Forgot Password ---")
		fmt.Print("Enter your email: ")
		inputEmail := readInput(reader)
		if _, exists := a.Users[inputEmail]; !exists {
			fmt.Print("Email not found, press enter to restart")
			readInput(reader)
			continue
		}

		fmt.Print("Enter a strong password: ")
		inputPassword := readInput(reader)

		fmt.Print("Confirm your password: ")
		inputConfirm := readInput(reader)

		if inputPassword != inputConfirm {
			fmt.Print("Wrong confirm password, press enter to back!")
			readInput(reader)
			continue
		} else {
			a.Users[inputEmail].Password = inputPassword
			fmt.Print("Password changed, press enter to back")
			readInput(reader)
			break
		}
	}
}
