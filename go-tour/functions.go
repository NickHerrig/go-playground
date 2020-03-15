package main

import (
    "fmt"
)

func basal_metabolic_rate(age, height, weight float64, sex bool) float64 {
    // weight in kg
    // height in cm
    // s = 5 for male and -161 for females

    s := 5.0
//    fmt.Println(reflect.TypeOf(s))
//    fmt.Println(reflect.TypeOf(age))
//    fmt.Println(reflect.TypeOf(height))
//    fmt.Println(reflect.TypeOf(weight))
    bmr := (10.0 * weight) + (6.25 * height) - (5.0 * age) + s
    return bmr
}

func main() {
    fmt.Println(basal_metabolic_rate(26, 167.64, 74.3891, true))
}
