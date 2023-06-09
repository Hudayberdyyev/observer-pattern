package objects

import "fmt"

type Customer struct {
	Id string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, itemName)
}

func (c *Customer) GetId() string {
	return c.Id
}
