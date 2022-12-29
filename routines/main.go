package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	for _, link:= range links {
		    go checkLink(link,c)
	}
	// here we are receivin the messages from the channels
// range c here is equal to <-c


   for l:=range c{
	go func(link string){
		time.Sleep(5*time.Second)
		 checkLink(link,c)
	}(l)
	
   }
}

func checkLink(link string,c chan string ) {
	time.Sleep(5*time.Second)
   _,err :=http.Get(link)

	if err!=nil{
  fmt.Println(link ,"might be down bad") 
   c <-link

}

  fmt.Println(link ,"is up")

  c<-link
}