package main

import "fmt"
type contactinfo struct{
	email string 
	zipcode int
}

type person struct {
	firstname string
	lastname  string
	 contactinfo
}

func main() {
	jim := person{ 
	firstname :"paul", 
	lastname:  "iwowo",
	  contactinfo:contactinfo{
		email:"iwowopaul66@gmail.com",
		zipcode:65,
		},

}

jim.updatefirstname("paullllo")
jim.print()


}

func (p person) print(){
	fmt.Printf("%+v",p)
 }

func ( pointertoperson *person) updatefirstname(name string){
	(*pointertoperson).firstname = name

}



