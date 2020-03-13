package main

import (
	"fmt"

	"github.com/acm-uiuc/core/controller"
	_ "github.com/acm-uiuc/core/database/migration"
	"github.com/acm-uiuc/core/service"
)

func main() {
	svc, err := service.New()
	if err != nil {
		fmt.Println(err)
	}

	controller, err := controller.New(svc)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(controller)
}
