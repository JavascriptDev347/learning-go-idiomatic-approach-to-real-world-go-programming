package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s, %s, %d", p.FirstName, p.LastName, p.Age)
}

//	p := Person{
// 	FirstName: "John",
// 	LastName:  "Doe",
// 	Age:       30,
// }

// fmt.Println(p.String())

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("Total: %d, Last Updated: %s", c.total, c.lastUpdated.Format(time.RFC3339))
}

type MailCategory int

const (
	Uncategoriezed MailCategory = iota + 90
	Personal
	Spam
	Social
	Advertisements
)

// embedding
type Inner struct {
	X int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.X * 2)
}

type Outer struct {
	S string
	X int
	Inner
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func main() {

	var c Counter
	c.Increment()
	fmt.Println(c.String())

	c2 := &Counter{}
	fmt.Println(c2.String())
	c2.Increment()
	fmt.Println(c2.String())

	fmt.Println("----------------------------")
	fmt.Println(Social)
	o := Outer{
		X: 10,
		Inner: Inner{
			X: 90,
		},
	}

	fmt.Println(o.X)       // 10
	fmt.Println(o.Inner.X) //90

	// embedin not inherit
	u := Outer{
		Inner: Inner{
			X: 20,
		},
		S: "Hello",
		X: 12,
	}

	fmt.Println(u.Double())

}
