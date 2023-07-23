package main

import "fmt"

var name string = "Sheikh Salah Uddin"

func main() {
    // var x float64 = 20.0
    // // Late Initialization
    // var num int
    // var amount float32
    // var isValid bool
    // var name string
    // num = 20
    // amount = 49.99
    // isValid = true
    // name = "Bappy"
    // fmt.Println(x)
    // fmt.Printf("x is of type %T\n", x)

    // var num = 20
    // var amount = 49.99
    // var isValid = true
    // var name = "Ruhul"
    // fmt.Println(num, amount, isValid, name)

    //Using the shorthand syntax
    // y := 33
    // fmt.Println(y)
    // fmt.Printf("y is of type %T\n",y)

    // Declare Multiple Variables
    // num1, num2 := 44, 45
    // amount, name := 48.23, "Sheikh" 
    // fmt.Println(num1, num2, amount, name)

    // var a, b, c = 1, 2, "sk"
    // fmt.Println(a)
    // fmt.Println(b)
    // fmt.Println(c)
    // fmt.Printf("a is of type %T\n", a)
    // fmt.Printf("b is of type %T\n", b)
    // fmt.Printf("c is of type %T\n", c)

    // var num int
    // var amount float32
    // var name string
    // var isValid bool
    // fmt.Println(num, amount, name, isValid)

    // variable redeclaration and reassigned
    // var_1, var_2 := 1, "hi"
    // fmt.Println(var_1, var_2)
    // var_3, var_2 := 2, "hello"
    // fmt.Println(var_3, var_2)
    // if var_4, var_2 := 3, "hey"; var_4 > var_1{
    //     fmt.Println(var_4,var_2)
    // }

    // fmt.Println(var_2)
    // fmt.Println(var_4)

    // var_5, var_2 := 3, 3.33
    // fmt.Println(var_5, var_2)
    // fmt.Println("Main My name is " + name)
    // printName()

    //Constant

    const myConst = 30
    fmt.Println(myConst)
    myConst = 20 // cant change myConst value
    fmt.Println(myConst)
}

// func printName(){
//     fmt.Println("My name is " + name)
// }