package main

import (
	"fmt"
)

func main() {
	dict := map[string] string {
		"0": "shashank",
		"1": "mohit",
	}

	fmt.Println(dict)

	dict["2"] = "shazam"

	fmt.Println(dict)
	
	delete(dict, "0")
	for key, value := range dict {
		fmt.Println(key, value)
	}
}