package main

import (
	"fmt"

	_ "github.com/acm-uiuc/core/database/migration"
	"github.com/acm-uiuc/core/server"
)

func main() {
	svr, err := server.New()
	if err != nil {
		fmt.Println(err)
	}

	svr.Start(":8000")
}
