package main

import (
	"awesomeProject/internal/api"
)

func main() {
	api := api.NewApi(api.Config{
		Address: ":9928",
	})

	api.Run()
}
