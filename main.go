package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	var tuidStr, locStr string

	switch len(os.Args) {
	case 2:
		tuidStr = os.Args[1]
		locStr = "Europe/Moscow"
	case 3:
		tuidStr = os.Args[1]
		locStr = os.Args[2]
	default:
		fmt.Println("Incorrect input, please try again")
		return
	}

	here, _ := time.LoadLocation(locStr)

	tuid, err := gocql.ParseUUID(tuidStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("time:", tuid.Time().In(here))
}
