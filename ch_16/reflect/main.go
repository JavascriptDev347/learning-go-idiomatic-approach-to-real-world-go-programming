package main

import (
	"fmt"
	"reflect"
)

type UserID int

func main() {
	var id UserID = 42

	t := reflect.TypeOf(id)
	fmt.Println(t.Name()) 
	fmt.Println(t.Kind())
}
