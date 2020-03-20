package main

import "fmt"

func halfs(num int) (int, bool) {
    return num/2, num%2==0
}

func greatest(nums ...int) int {
    var max int
    for num := range nums {
        if num > max {
            max = num
        }
    }
    return max 
}

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func main() {
    nextodd:= makeOddGenerator()
    fmt.Println(nextodd())
    fmt.Println(nextodd())
    fmt.Println(nextodd())
    fmt.Println(nextodd())
}
