// main.go
package main

import (
	"fmt"
	"log"

	"github.com/ssebs/padpal-server/api"
)

func main() {
	fmt.Println("PadPal Server")
	log.Fatal(api.DoEverything("0.0.0.0:5000"))
}
