package main

import (
	"abelce/app/gateway/cmds"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Printf("%v", os.Args)
	log.SetFlags(log.Lshortfile)
	app := cmds.NewCliApp()
	app.Run(os.Args)
}
