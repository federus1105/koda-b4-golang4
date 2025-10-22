package internals

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type AuthInterface interface {
	Register(firstname, lastname, email, password, confirm string)
	Login(email, password string)
}

type AuthSystem struct {
	Users map[string]*User
}

func readInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (a *AuthSystem) Register() {
	var inputemail string
	var inputFirstname string
	var inputLastname string
	var inputpassword string
	var inputconfirm string
	// var Continue string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- Register ---")
	fmt.Print("What is your first name: ")
	inputFirstname = readInput(reader)
	fmt.Print("What is your last name: ")
	inputLastname = readInput(reader)
	fmt.Print("What is your email: ")
	inputemail = readInput(reader)
	fmt.Print("Enter a strong password: ")
	inputpassword = readInput(reader)
	fmt.Print("Confirm your password: ")
	inputconfirm = readInput(reader)
	if inputpassword != inputconfirm {
		panic("Wrong confirm password, press enter to back!")
	}
	a.Users[inputemail] = &User{
		Firstname: inputFirstname,
		Lastname:  inputLastname,
		Email:     inputemail,
		Password:  inputpassword,
	}
	fmt.Println("\nIs it true?")
	fmt.Println("firstname : ", inputFirstname)
	fmt.Println("lastname : ", inputLastname)
	fmt.Println("email : ", inputemail)
	fmt.Print("Continue? (Y/n): ")
	// Continue := readInput(reader)
	// fmt.Print(Continue)

	// var choice string
	// fmt.Scanln(&choice)
	// switch choice {
	// case "Y":
	// 	fmt.Print(Continue)
	// case "N":

	// }
}

func (a *AuthSystem) Login() {
	var inputemail string
	var inputpassword string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--- LOGIN ---")
	fmt.Print("Masukkan Email: ")
	inputemail = readInput(reader)
	fmt.Print("Masukkan Password: ")
	inputpassword = readInput(reader)
	user, exists := a.Users[inputemail]
	if !exists || user.Password != inputpassword {
		panic("Wrong email or password, press enter to restart")
	}
	fmt.Println("Login success, press enter to back..")
}

