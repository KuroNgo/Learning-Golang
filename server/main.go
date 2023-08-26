package main

import (
	"server/route"
)

func main() {
	app := route.NewService()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
