// main.go
package main

import (
	"fmt"
	"log"

	"github.com/ssebs/padpal-server/api"
)

func main() {
	fmt.Println("PadPal Server")
	a := api.NewAPI("0.0.0.0:5000", "./tmp/")
	log.Fatal(a.RunAPI())
}
