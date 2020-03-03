package main

import (
	"fmt"
	"flag"

	"github.com/acm-uiuc/core/database/migrations"

	_ "github.com/acm-uiuc/core/services/auth"
	_ "github.com/acm-uiuc/core/services/user"
)

type cliFlags struct {
	migration string
}

func main() {
	flags := cliFlags{}
	flag.StringVar(&flags.migration, "migration", "", "the migration to start from")
	flag.Parse()

	if flags.migration != "" {
		err := migrations.Migrate(flags.migration)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		return
	}
}
