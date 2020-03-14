package main

import (
	"flag"
	"fmt"

	"github.com/acm-uiuc/core/database/migration"
	"github.com/acm-uiuc/core/server"
)

type options struct {
	migration string
	server    bool
}

func main() {
	opts := options{}
	flag.StringVar(&opts.migration, "migration", "", "migration to start from")
	flag.BoolVar(&opts.server, "server", false, "enable to run the server")
	flag.Parse()

	if opts.migration != "" {
		err := migration.Migrate(opts.migration)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		return
	}

	if opts.server {
		svr, err := server.New()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		err = svr.Start(":8000")
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		return
	}
}
