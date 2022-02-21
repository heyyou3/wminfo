package main

import (
	"fmt"
	"log"
	"wminfo/window"
)

func main() {
	winClient, err := window.New()
	if err != nil {
		log.Fatal(err)
	}

	windows, err := winClient.FetchWindowInfo()
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range windows {
		fmt.Printf("ID = %d\n", w.ID)
		fmt.Printf("Name = %s\n", w.Name)
		fmt.Printf("Class = %s\n", w.WmClass.Class)
	}
}
