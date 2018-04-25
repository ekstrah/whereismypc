package main

import (
	"fmt"

	"github.com/lextoumbourou/goodhosts"
)

func main() {
	// editGoodHosts()
	ViewGoodHosts()
}

func ViewGoodHosts() {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		panic(err)
	}

	for _, line := range hosts.Lines {
		fmt.Println(line.Raw)
	}
}

func editGoodHosts() {
	hosts, err := goodhosts.NewHosts()
	if err != nil {
		panic(err)
	}

	hosts.Add("127.0.0.1", "searchmyip.com")

	if err := hosts.Flush(); err != nil {
		panic(err)
	}
}
