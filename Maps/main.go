package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// colors := make(map[string]string)

	delete(colors, "green")

	// fmt.Println(colors)
	updateMap(colors, "red", "F***")
	printMap(colors)
}

func updateMap(c map[string]string, key string, value string) {
	c[key] = value
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex of", color, "is", hex)
	}
}
