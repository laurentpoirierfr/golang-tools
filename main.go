package main

import (
	"log"

	"github.com/laurentpoirierfr/golang-tools/profile"
)

func main() {

	value := profile.GetValueString("application.profile")
	log.Println("application.profile :", value)

}
