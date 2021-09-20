package main

import "fmt"

type Customer struct {
	FirstName, LastName, Email string
	Address                    []Address
}

type Address struct {
	Street, City, State, Zip string
}

func (c Customer) ToString() string {
	return fmt.Sprintf("Customer: %s %s, Email:%s", c.FirstName, c.LastName, c.Email)
}

func (c Customer) ShippingAddress() string {
	for _, v := range c.Address {

		return fmt.Sprintf("%s, %s, %s, Zip - %s", v.Street, v.City, v.State,
			v.Zip)

	}
	return ""
}
func (c *Customer) ChangeEmail(newEmail string) {
	c.Email = newEmail
}
func main() {

	var c Customer
	c.FirstName = "Mainul"
	c.LastName = "Hasan"
	c.Email = "mainul@rec.com"

	Address := []Address{
		{
			Street: "Gollamari",
			City:   "BaniyaKhamar",
			State:  "Khulna",
			Zip:    "1290",
		},
		{
			Street: "Bhuapur",
			City:   "Tangail",
			State:  "Dhaka",
			Zip:    "1960",
		},
	}
	c.Address = Address
	c.ChangeEmail("test@emal.com")
	fmt.Println(c.ToString())
	fmt.Println(c.ShippingAddress())

}
