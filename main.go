package main

import (
	"bmg/configs"
	"bmg/routes"
)

func main() {
	configs.InitDB()

	e := routes.New()
	e.Start(":8000")
}