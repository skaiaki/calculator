package main

import (
	"fmt"
)

func add(a,b float64)(float64){
    return a+b
}

func subtract(a,b float64)(float64){
    return a-b
}

func multiply(a,b float64)(float64){
    return a*b
}
func divide(a,b float64)(float64){
    return a/b
}

func main() {
    a, b, s := getInput()
    var res float64
    switch s {
    case "+":
        res=add(a,b)

    case "-":
        res=subtract(a,b)  
    
    case "*":
        res=multiply(a,b)
    
    case "/":
        res = divide(a, b) 
        
    default:
        fmt.Scanln("wrong sign")  
     return     
    }
    fmt.Printf("Результат: %.2f\n", res)  
}

func getInput ()(float64, float64, string){
var a, b float64
var s string

fmt.Print ("enter the first number: ")
fmt.Scanln(&a)

fmt.Print ("enter the second number: ")
fmt.Scanln(&b)

fmt.Print ("enter the second operator: ")
fmt.Scanln(&s)

return a, b, s
}
