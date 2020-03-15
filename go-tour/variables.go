package main

import (
    "fmt"
    "reflect"
)

// variables with initalizers
var protien int = 200
const carb = 150
var fat string = "Yes!"

func main() {
fmt.Println(protien, fat, carb)
fmt.Println(reflect.TypeOf(protien), reflect.TypeOf(fat), reflect.TypeOf(carb))

// if an initalizer is present, no need for a type declaration...
var micronutrient = "Yes!"
fat:= "No!"
protien := float64(protien)

fmt.Println(protien, fat, carb, micronutrient)
fmt.Println(reflect.TypeOf(protien), reflect.TypeOf(fat), reflect.TypeOf(carb), reflect.TypeOf(micronutrient))
}
