package main

 import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	
 )

type deck []string

func newdeck() deck {
	cardsbody := deck{}
	cardsuits := []string{"spades", "king", "hearts", "diamonds"}
	cardvalues := []string{"Ace", "two", "three", "four"}

	for _, cards := range cardsuits {
		for _, cardvalues := range cardvalues {
			cardsbody = append(cardsbody, cardvalues+"of"+cards)
		}
	}

	return cardsbody
}
// reciever function
func (d deck) print() {

	for i,cards := range d {
		fmt.Println(i,cards)
	}  

}
// normal function

func deal( deal deck ,handsize int)(deck,deck) {
	  return deal[:handsize],deal[handsize:]
	
}
// reciever function to turn a  slice of string to a single string

// this functions change the deck from a slice of strings to a string  
func(d deck) tostring() string{
   joined:=[]string(d)
  return strings.Join(joined, ",")

}

//this functions changes the string from a byte of slice to a 

func(d deck) savetofile(filename string) error {
	
	 return ioutil.WriteFile(filename,[]byte(d.tostring()),0666)
}

func createdeckfromfile(filename string)  deck {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("File contents: %s", content)
   s:=strings.Split(string(content),",")
   return deck(s)

}



func( d deck) shuffledeck(){
	for index := range d {
  newposition:= rand.Intn(len(d))
   d[index],d[newposition] = d[newposition],d[index]

	}
}


// my own functions 
