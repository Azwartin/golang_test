package main

import (
	"fmt"
	"os"
	"tagvalidator/readers"
	"tagvalidator/validators"
)

func main() {
	args := os.Args[1:] //get args without program

	if len(args) < 1 {
		fmt.Println("Filename or url required,for example tagvalidator /var/www/html.html")
		return
	}

	path := args[0]
	reader := readers.CreateReader(path)
	html, err := reader.Read(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	validator := validators.Validator{}

	if validator.Validate(html) {
		fmt.Println("Tags are correct")
	} else {
		fmt.Println("Tags are incorrect")
	}
}
