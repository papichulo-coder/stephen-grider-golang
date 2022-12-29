package main

import "fmt"

func main() {
	color := map[string]string{
		"red":   "johnn",
		"green": "paul",
		"blue":  "jolla",
	}
	// update a value  in  a  map in golang
	// colors["red"] = 55
	// // delete a value in a map
	// delete(colors,"red")
	// //fmt.Println(colors["red"]) 
	// fmt.Println(colors)
printMap(color)

}


func printMap( c map[string]string)  {
   for colors,hex := range c{
     fmt.Println("hex code for ", colors,"is", hex)
   }
}