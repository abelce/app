package main

import (
	"app/graphql/cmds"
	"fmt"
	"os"
	// _ "github.com/lib/pq"
)

func main() {
	fmt.Printf("%v", os.Args)
	app := cmds.NewCliApp()
	app.Run(os.Args)
}
