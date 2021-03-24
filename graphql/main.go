package main

import (
	"fmt"
	"os"
	"vwood/app/graphql/cmds"
	// _ "github.com/lib/pq"
)

func main() {
	fmt.Printf("%v", os.Args)
	app := cmds.NewCliApp()
	app.Run(os.Args)
}
