package main

import "observer-pattern/objects"

func main() {
	shirtItem := objects.NewItem("Nike Shirt")

	observerFirst := &objects.Customer{Id: "abc@gmail.com"}
	observerSecond := &objects.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
}
