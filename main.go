package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type struct 

func main(){
	args := os.Args[1:]

	switch len(args) {
	case 1:
		fmt.Println("Size of json: ", args[0], "kb")
		createJSON()
	default:
		fmt.Printf("Incorrect number of args")
	}
}

func createJSON(){
	
}