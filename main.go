package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocql/gocql"
)

const defaultLocation = "Etc/GMT"

func main() {
	var tuidStr, locStr string

	switch len(os.Args) {
	case 2:
		tuidStr = os.Args[1]
		locStr = defaultLocation
	case 3:
		tuidStr = os.Args[1]
		locStr = os.Args[2]
	default:
		fmt.Println("Incorrect input, please try again")
		return
	}

	here, err := time.LoadLocation(locStr)
	if err != nil {
		fmt.Println("Can't parse location, please try again")
		return
	}

	tuid, err := gocql.ParseUUID(tuidStr)
	if err != nil {
		fmt.Println("Can't parse timeuuid, please try again")
		return
	}

	fmt.Println("time:", tuid.Time().In(here))
}
