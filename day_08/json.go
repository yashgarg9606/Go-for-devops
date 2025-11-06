package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{"Yash", 19}

	jsonData, err := json.Marshal(person) //Marshal is serialization of data
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON Data:", string(jsonData))

	var person2 Person

	newErr := json.Unmarshal(jsonData, &person2) //Unmarshal is deserialization of data
	if newErr != nil {
		log.Fatal(newErr)
	}

	fmt.Printf("Unmarshaled Data: %+v\n", person2)
}
