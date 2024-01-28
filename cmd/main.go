// main.go
package main

import (
	"fmt"

	"github.com/ssebs/padpal-server/api"
)

func main() {
	fmt.Println("PadPal Server")

	api.HandleAndServe("", 5000)
}
