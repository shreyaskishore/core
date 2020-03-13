package main

import (
	"fmt"

	_ "github.com/acm-uiuc/core/database"
	_ "github.com/acm-uiuc/core/database/migration"
	_ "github.com/acm-uiuc/core/model"
	_ "github.com/acm-uiuc/core/service"
)

func main() {
	fmt.Println("Hello World!")
}
