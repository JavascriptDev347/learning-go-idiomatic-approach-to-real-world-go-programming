package main

import "encoding/json"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	user := User{
		Name: "John Doe",
		Age:  30,
	}

	// marshal
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	println(string(data))

	jsonData := []byte(`{"name":"Jane Doe 2","age":25}`)

	var user2 User

	err = json.Unmarshal(jsonData, &user2)
	if err != nil {
		panic(err)
	}

	println(user.Name, user.Age)

}
