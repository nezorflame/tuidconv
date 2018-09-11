package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
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
		fmt.Println("Incorrect input")
		return
	}

	here, err := time.LoadLocation(locStr)
	if err != nil {
		fmt.Printf("Unable to parse location: %v\n", err)
		return
	}

	tuid, err := uuid.Parse(tuidStr)
	if err != nil {
		fmt.Printf("Unable to parse timeuuid: %v\n", err)
		return
	}

	sec, nsec := tuid.Time().UnixTime()
	fmt.Println("time:", time.Unix(sec, nsec).In(here))
	fmt.Println("unix:", sec)
}
