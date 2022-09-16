package main

import "demoProject/src/AHP_Gin/routers"

func main() {
	r := routers.Routers()
	r.Run(":9090")
}
