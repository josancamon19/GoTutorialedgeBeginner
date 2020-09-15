package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

func main() {
	// Json marshalling
	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}
	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

	// Json unmarshall
	var bookUnmarshalled Book
	err = json.Unmarshal(byteArray, &bookUnmarshalled)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bookUnmarshalled)

	// Unstructured data
	var bookMap map[string]interface{}
	err = json.Unmarshal(byteArray, &bookMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bookMap)
}
