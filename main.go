package main

import(
	"fmt"
	db "app/database"
)

func main() {

	s := db.GetBase()

	if s[0].Status{
		fmt.Println("Ishladi")
	} else {
		fmt.Println("Ishlamadi")
	}


}