package main

import (
	"log"

	"github.com/laurentpoirierfr/golang-tools/config"
)

func main() {

	value := config.GetIntegerValue("application.port")
	log.Println("application.port :", value)

}
