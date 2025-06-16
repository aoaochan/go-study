package t02greetings

import "fmt"

func Hello(name string) string {
	var message string = fmt.Sprintf("Hi, &v. Welcome!", name)
	return message
}
