package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	// Marshal va Encoder farqi
	user := User{
		Name: "Rustambek",
		Age:  25,
	}

	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	user2 := User{
		Name: "Rustambek",
		Age:  221,
	}
	//Bu yerda NewEncoder() ga os.Stdout berdik.
	// os.Stdout — terminalga yozish imkonini beruvchi io.Writer.
	// Encoder ma'lumotni to'g'ridan-to'g'ri shu joyga yozadi.
	// Orada alohida []byte natijasini olishimiz shart emas.
	err = json.NewEncoder(os.Stdout).Encode(user2)
	if err != nil {
		log.Fatal(err)
	}

	jsonData := []byte(`{"name":"Rustambek","age":25}`)

	var user3 User

	err = json.Unmarshal(jsonData, &user3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
