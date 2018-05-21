package main

import "fmt"

func main() {
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}
	var colors2 map[string]string
	colors3 := make(map[string]string)
	printMap(colors1)
	printMap(colors2)
	printMap(colors3)
}

func printMap(c map[string]string) {
	for color, hexColor := range c {
		fmt.Println("Hex code for", color, "is", hexColor)
	}
}
