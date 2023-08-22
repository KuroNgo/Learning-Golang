package greetings

import (
	"fmt"
)

func Hello(name string) string {
	//Return a greetings that embeds the name in  a message
	// Sprintf cho người dùng nhập vào
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}
