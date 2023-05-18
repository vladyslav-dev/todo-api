package main

import (
	"github.com/vladyslav-dev/todo-api/app"
)

func main() {

	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
