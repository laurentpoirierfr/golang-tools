package main

import (
	"log"

	"github.com/laurentpoirierfr/golang-tools/profile"
)

func main() {

	value := profile.GetIntegerValue("application.port")
	log.Println("application.port :", value)

}
