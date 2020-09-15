package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func main5() {
	parseJson()
}

func parseJson() {
	jsonFile, err := os.Open("files/users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Opened Users.json")
	defer jsonFile.Close()

	jsonData, _ := ioutil.ReadAll(jsonFile)
	var users Users
	// we unmarshal our byteArray(jsonData) which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(jsonData, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Printf("User is %s \n", users.Users[i].Name)
	}
}
