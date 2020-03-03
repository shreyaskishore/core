package main

import (
	"fmt"
	"flag"

	"github.com/acm-uiuc/core/database/migrations"

	"github.com/acm-uiuc/core/services"
)

type cliFlags struct {
	migration string
	server bool
}

func main() {
	flags := cliFlags{}
	flag.StringVar(&flags.migration, "migration", "", "the migration to start from")
	flag.BoolVar(&flags.server, "server", false, "enable to run the server")
	flag.Parse()

	if flags.migration != "" {
		err := migrations.Migrate(flags.migration)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		return
	}

	if flags.server {
		_, err := services.New()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		// TODO: Bind controller and endpoints to services

		return
	}
}
