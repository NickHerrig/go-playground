package main

import (
    "fmt"
)

func main() {
      // Example of a runtime error!
//    var my_map map[string]int
//    my_map["key"] = 11
//    fmt.Println(my_map)
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["O"] = "Oxygen"
    fmt.Println(elements)

// or a better way...
    greetings := map[string]string{
        "English": "Hello World",
        "Spanish": "Hola Mundo",
        "Italian": "Ciao Mondo",
        "French" : "Bonjour Le Monde",
    }
    fmt.Println(greetings)

    // Iterationg over key and value
    for key, value := range elements {
        fmt.Println("The element symbol:",key,"stands for", value)
    }

    // looking up an non-existant element
    name, present := elements["Nick"]
    fmt.Println(name, present)
}
