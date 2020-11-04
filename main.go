package main

import (
	"github.com/laurentpoirierfr/golang-tools/config"
	"github.com/laurentpoirierfr/golang-tools/log"
)

func main() {

	value := config.GetStringValue("application.port")
	log.Debug("application.port : " + value)

}
