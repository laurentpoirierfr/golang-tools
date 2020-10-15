package tools

import (
	"fmt"
	"log"
)

// FailOnError gestion des erreurs
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
