package main

import (
	"fmt"
	"regexp"
)

func main() {

	var validEmail, _ = regexp.Compile("^[_a-zA-Z0-9+-.]+@[a-zA-Z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$")

	fmt.Println(validEmail.MatchString("abc@domain.com"))                 // true
	fmt.Println(validEmail.MatchString("abc@domain.io"))                  // true
	fmt.Println(validEmail.MatchString("abc62@abc-domain.com"))           // true
	fmt.Println(validEmail.MatchString("abc62@subdomain.abc-domain.com")) // true

	fmt.Println(validEmail.MatchString("abc62@abc-domain")) // flase
	fmt.Println(validEmail.MatchString("abc@abc.co.krrrr")) // false

}
