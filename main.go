package main

import (
	"table_management/api"
)

func main() {
	api.NewApiServer().Run()
}
