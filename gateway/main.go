package main

import (
	"fmt"
	"log"
	"os"
	"vwood/app/gateway/cmds"
)

func main() {
	fmt.Printf("%v", os.Args)
	log.SetFlags(log.Lshortfile)
	app := cmds.NewCliApp()
	app.Run(os.Args)
}
