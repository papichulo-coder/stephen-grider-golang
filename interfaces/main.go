package main

import "fmt"

type bot interface{
	getGreeting() string

}

type englishbot struct {
}

type spanishbot struct {
}

func main() {
eb:=englishbot{}
sb:=spanishbot{}
printgreeting(eb)
printgreeting(sb)
}

func printgreeting(b bot){
  fmt.Println(b.getGreeting())
}
func ( englishbot) getGreeting() string {
	return "awfa my guy"
}

func (spanishbot) getGreeting() string {
	return "hola amigos"
}

