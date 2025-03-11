package main

import "fmt"

//define a struct

type car struct{
	Brand string
	colour string
}
type byke struct{
	make string
	chasis string
	age string

}

type house struct{
	make string
	age string
	size string
	rooms string
}


//define a method for the  car struct
func (c car) Drive(){
	fmt.Println(c.colour ,c.Brand , "is driving")

}
func(b byke) ride(){
	fmt.Println(b.chasis, b.make, "is more reliable compared to an", b.age, "ferrari")
}
func(h house) built(){
	fmt.Println("The Ambani's",h.age, h.rooms, h.make,"house is quite", h.size)
}
func main(){
	myCarr := car{Brand: "Bentley", colour: "Emerald green"}

	myCarr.Drive()
	
	mybyke := byke{chasis:"manual", age:"old", make:"Ducati"}
	mybyke.ride()

	myhouse := house{make:"bungalow", age:"new", size:"huge", rooms:"3 bedroom"}
	myhouse.built()
}

