// map -> pythonのdictと同じ

package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4b7f45",
		"white": "#ffffff",
	}

	// colors := make(map[string]string) // make -> 空のデータ構造を作る

	// delete(colors, "red") // delete key

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) { // iteration key とvalueが渡される
	for color, hex := range c {
		fmt.Println("Hex code for ", color, "is", hex)
	}
}
