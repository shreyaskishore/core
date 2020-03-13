package main

import (
	"fmt"

	_ "github.com/acm-uiuc/core/model"
	_ "github.com/acm-uiuc/core/service/auth"
	_ "github.com/acm-uiuc/core/service/group"
	_ "github.com/acm-uiuc/core/service/user"
)

func main() {
	fmt.Println("Hello World!")
}
