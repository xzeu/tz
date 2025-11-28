package main

import (
	"fmt"

	"github.com/xzeu/tz/cmd/tz"
)

func main() {
	err := tz.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
