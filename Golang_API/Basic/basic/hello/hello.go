package main

import (
	"fmt"

	"example/greetings"

	"log"
)

func main() {
	// Thiết lập thuộc tính của predefind logger, bao gồm
	// Log entry prefix và một flaf để disable print thười gian,
	// nguồn file và số dòng
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//A slice of names
	names := []string{"Kuro", "Kaze", "Phong"}
	// Request a greetings mesage
	message, err := greetings.Hellos(names)
	//Request greeting mesage gỏ the names
	// Nếu một lỗi được trả về, in nó ra console và thoát khỏi chương trình
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
