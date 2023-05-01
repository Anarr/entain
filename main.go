package main

import (
	"github.com/Anarr/entain/cmd"
	_ "github.com/Anarr/entain/docs"
)

// @title Entain Request Process Service
// @version 0.2.13
// @description Entain handle user requests and process them

// @contact.name Anar
// @contact.email anar.rzayev94@gmail.com

// @schemes http
// @BasePath /
func main() {
	cmd.Execute()
}
