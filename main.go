package main

import (
	"fmt"
	"flag"

	"github.com/acm-uiuc/core/database/migrations"

	"github.com/acm-uiuc/core/services"
	"github.com/acm-uiuc/core/controllers"
)

type cliFlags struct {
	migration string
	server bool
	site string
}

func main() {
	flags := cliFlags{}
	flag.StringVar(&flags.migration, "migration", "", "migration to start from")
	flag.BoolVar(&flags.server, "server", false, "enable to run the server")
	flag.StringVar(&flags.site, "site", "", "path to serve site from")
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
		svcs, err := services.New()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		controller := controllers.New(svcs, flags.site)
		controller.Start(":8000")
		return
	}
}
