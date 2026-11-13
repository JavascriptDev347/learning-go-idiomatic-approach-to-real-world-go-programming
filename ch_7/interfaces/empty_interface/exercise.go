package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func getData() {

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	fmt.Println("Json file successfully opened!!!")
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var users Users

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for _, user := range users.Users {
		fmt.Println("User Type: " + user.Type)
		fmt.Println("User Age: " + strconv.Itoa(user.Age))
		fmt.Println("User Name: " + user.Name)
		fmt.Println("Facebook Url: " + user.Social.Facebook)
	}

}

func main() {
	getData()
}
