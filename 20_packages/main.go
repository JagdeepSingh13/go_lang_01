package main

import (
	"fmt"

	"github.com/JagdeepSingh13/go_lang_01/auth"
	"github.com/JagdeepSingh13/go_lang_01/user"
	"github.com/fatih/color"
)

func main() {
	// if name starts with capital-letter then can be used anywhere

	auth.LoginWithCredentials("Jsingh", "secret")

	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "j@j.com",
		Name:  "JSingh",
	}
	fmt.Println(user)

	color.Blue(user.Name)

}
