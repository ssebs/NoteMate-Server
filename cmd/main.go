package main

import (
	"fmt"

	"github.com/ssebs/padpal-server/api"
)

func main() {
	fmt.Println("PadPal Server")

	fmt.Println(api.APIFoo())
}
