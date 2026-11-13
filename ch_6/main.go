package main

import "encoding/json"

func main() {
	f := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}

	err := json.Unmarshal([]byte(`{"name":"Bob","age":30}`), &f)
	if err != nil {
		panic(err)
	}
	println(f.Name, f.Age) // prints Bob 30
}
