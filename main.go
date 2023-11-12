package main

import "go_ranking/router"

func main() {
	r := router.Router()
	r.Run(":9999")
}
