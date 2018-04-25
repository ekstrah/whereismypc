package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	editHost()
}

func editHost() {
	f, err := os.OpenFile("C:/Windows/System32/drivers/etc/hosts", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	len, err := f.WriteString("127.0.0.1 ekstrah.xyz")
	if err != nil {
		log.Fatalf("failed writing the file", err)
	}
	fmt.Println(len)
}
