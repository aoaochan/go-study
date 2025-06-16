package main

import (
	"fmt"
	"log"

	"aoaochan/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var names = []string{"aaa", "bbb", "ccc"}
	var messages, err = greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		fmt.Printf("%v: %v\n", name, message)
	}

	// var message, err = greetings.Hello("aoaochan")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)
}
