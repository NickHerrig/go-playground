package main

import (
    "fmt"
)

func basal_metabolic_rate(age, height, weight float64, sex string) float64 {
    // weight in kg
    // height in cm
    // s = 5 for male and -161 for females
    if sex == "male"{
        bmr := (10.0 * weight) + (6.25 * height) - (5.0 * age) + 5.0
        return bmr
    } else if sex == "female" {
        bmr := (10.0 * weight) + (6.25 * height) - (5.0 * age) - 161.0
        return bmr
    }
    return 0
}

func main() {
    fmt.Println(basal_metabolic_rate(26, 167.64, 74.3891, "male"))
    fmt.Println(basal_metabolic_rate(26, 167.64, 74.3891, "female"))
}
