package main

import (
	"hw1/internal/app"
	"log"
	"os"
)

//Явно используется контейнер, так как это требуется в задании 
func main() {
	container := app.NewContainer(os.Stdin, os.Stdout)

	if err := container.Runner.Run(); err != nil {
		log.Fatal(err)
	}
}
