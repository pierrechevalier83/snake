package main
import (
    "fmt"
)

func main() {
	fmt.Println("This function sums A,B,C")
    fmt.Println("How much is A?")
    var A int
    fmt.Scanf("%v", &A)
    fmt.Println("A=", A)
    fmt.Println("How much is B?")
    var B int
    fmt.Scanf("%v", &B)
    fmt.Println("B=", B)
	fmt.Println("How much is C?")
	var C int
    fmt.Scanf("%v", &C)
    fmt.Println("C=", C)
    sum(A,B,C)
}

func sum(input ...int) int {
    sum := 0

    for i := range input {
        sum += input[i]
    }
	fmt.Println("")
    fmt.Println("A+B+C=", sum)
    return sum
}
	
