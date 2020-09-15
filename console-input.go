package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	readText()
	readChar()
}
func readText() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		fmt.Printf("Text entered is %s", text)
	}
}

func readChar() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	// print out the unicode value i.e. A -> 65, a -> 97
	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A Key Pressed")
		break
	case 'a':
		fmt.Println("a Key Pressed")
		break
	}
}
