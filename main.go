package main

import (
	"fmt"

	_ "github.com/acm-uiuc/core/database/migration"
	"github.com/acm-uiuc/core/service"
)

func main() {
	svc, err := service.New()
	fmt.Println(err)
	fmt.Println(svc)
}
