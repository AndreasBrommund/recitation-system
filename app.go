package main

import (
	"flag"
	"fmt"

	"github.com/DavidSkeppstedt/recitation/web"
)

func main() {
	port := flag.String("port", ":8080", "The port the webserver should run on.")
	flag.Parse()
	fmt.Println("Port:", *port)
	web.StartServer(*port)
}
