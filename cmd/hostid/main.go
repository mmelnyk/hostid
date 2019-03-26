package main

import (
	"fmt"

	"go.melnyk.org/hostid"
)

func main() {
	if id, err := hostid.ID(); err == nil {
		fmt.Println("Host ID:", id)
		return
	}
	fmt.Println("ERROR: Cannot get unique Host ID")
}
