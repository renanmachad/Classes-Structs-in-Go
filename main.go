package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	FirstName string
	LastName  string
	Years     int
}

func New(firstName string, lastName string, years int) Person {
	new_employee := Person{
		FirstName: firstName,
		LastName:  lastName,
		Years:     years,
	}
	return new_employee
}

func InputString(prompt string, r *bufio.Reader) (string, error) {
	fmt.Printf(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func main() {

	// create the class reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Create the person:")

	// create the functions to get the input of users
	name, _ := InputString("Name of person:", reader)
	surname, _ := InputString("Surname of person:", reader)
	years, _ := InputString("Years of person:", reader)

	// convert string to int
	years_int, err := strconv.Atoi(years)

	// error handler
	if err != nil {
		fmt.Println("ERROR")
		panic(err)
	}

	// create the struct/classe person
	e := New(name, surname, years_int)

	// Print the result
	fmt.Println(e.FirstName, e.LastName, e.Years)

}
