package main

import "fmt"

//defining a struct

type Toy struct{
	Name string
	color string
	Size string
}

type Car struct{
	Name string
	color string
	Size string
	form string
}

func main(){
	//creating a new toy (like making a new box)
	myToy := Toy {
		Name : "Teddy",
		color: "White",
		Size: "Large",
	}

	myCar := Car{
		Name : "Lamborgini Urus",
		color: "Emerald Green",
		Size:"huge car",
		form: "Manual car",


	}
	




//Using the toys info
    fmt.Println("My toy is", myToy.Name)
	fmt.Println("It is", myToy.color)
	fmt.Println("its", myToy.Size)


	//using cars info
	fmt.Println("My car is", myCar.Name)
	fmt.Println("It is", myCar.color)
	fmt.Println("its a",myCar.Size)
	fmt.Println("its a",myCar.form)
}

