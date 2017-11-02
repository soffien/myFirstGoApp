package main

import (
	"fmt"

	"github.com/soffien/myFirstGoApp/resources"
)

func main() {
	sof := resources.NewPerson("Sof", "ain sultan", 35)
	fmt.Printf("Sof content %#v \n", sof)
	err := sof.ChangeName("mo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("New Name %s\n", sof.GetName())

	newAge := sof.AddAge(5)
	fmt.Printf("New age %d\n", newAge)

	fmt.Printf("New age from object %d\n", sof.GetAge())
}
