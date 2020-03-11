package main

import "fmt"

// how to comment in go
func main() {
    var bmr float32 = 1667
    var activity_multiplier float32  = 1.5
    var total_cal float32 = bmr * activity_multiplier
    fmt.Println("Total Calories: ", total_cal)

    var weight float32 = 164
    var body_fat_percent float32 = 0.16
    var lean_body_mass float32 = weight * ( 1.0 - body_fat_percent )
    fmt.Println("Lean Body Mass:", lean_body_mass)
    var protien_rate float32 = 1.5
    var protien_macro float32 = protien_rate *  lean_body_mass 
    fmt.Println("Protien Macro: ", protien_macro)
}
