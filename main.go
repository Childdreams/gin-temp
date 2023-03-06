package main

import (
	_ "app/databases"
	"app/routers"
)

func main() {

	r := routers.Load()
	r.Run()
}
