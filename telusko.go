// Printing Hello world
// package main

// import "fmt"

// func main(){
// 	fmt.Println("Hello world")
// }


// Variables declaration and assignment
// package main

// import "fmt"

// func main(){
	// fmt.Println("The sum of 5 and 6 is :",5+6)

	// var num1 = 20
	// var num2 = 30
	// var result = num1+num2
	// fmt.Println("The sum of num1 and num2 is :",result)

	// var num1 int = 20
	// var num2 int = 30
	// var result int = num1+num2
	// fmt.Println("The sum of num1 and num2 is :",result)

	// var num1, num2, result int
	// num1 =20
	// num2 =30
	// result =num1+num2
	// fmt.Println("The sum of num1 and num2 is :",result)

	// var num1, num2 = 20, 30
	// var result =num1+num2
	// fmt.Println("The result is :",result)

	// num := 50
	// fmt.Println(num)

	// We can re-assign the variable again, it will store newest assigned value
	// num := 50
	// num = 40
	// fmt.Println(num)

	// We can not re-assign the const again, if we will try to re-assign, it will give error
	// const num =50
	// num = 40        //Error
	// fmt.Println(num)
// }



// loops(go has only for loop)

// package main

// import "fmt"

// func main(){


	// Infinite for loop

	// for {
	// 	fmt.Println("Infinite loop")
	// }


	//  1st way to define for loop
	// i :=1
	// for i<=5 {
	// 	fmt.Println("1st way, for loop", i)
	// 	i++
	// }


	// 2nd way to define for loop

	// for i :=1; i<=5; i++{
	// 	fmt.Println("2nd way, for loop", i)
	// }
// }




// functions creation
// package main

// import "fmt"

// func add(num1, num2 int) int{
// 	result :=num1 +num2

// 	return result
// }

// func main(){

// 	num1 := 30
// 	num2 :=20
// 	result := add(num1, num2)

// 	fmt.Println("The sum of num1 and num2 is :",result)
// }



// function returning multiple things
// package main

// import "fmt"

// func calc(num1, num2 int) (int, int){
// 	add := num1 + num2
// 	subs := num1 - num2

// 	return add, subs
// }

// ------------------or-----------------------

// func calc(num1, num2 int)(add, subs int){
// 	add = num1 + num2
// 	subs = num1 - num2
	
// 	return
// }

// func main(){

// 	num1 := 30
// 	num2 :=20
// 	add, subs := calc(num1, num2)

// 	fmt.Println("The sum of num1 and num2 is :",add)
// 	fmt.Println("The difference between num1 and num2 is :",subs)
// }






// variable scope
// function level(declaring inside function )
		// func demo{
		// 	a :=10   // a is only accessed inside function demo only
		// }

// Package level variable(declaring outside function "var a = 10")(accecable to the whole package)
// package main

// import "fmt"

// var a = 10

// func demo(){
// 	fmt.Println("in demo",a)
// }

// func main(){
// 	fmt.Println("in main",a)
// 	demo()
// }


// declaring variables with same name at package and local level(it gives preference to local, 
// 											if local variable not present then it will take package variable)
// package main

// import "fmt"

// var a = 10

// func demo(){
// 	var a = 8
// 	fmt.Println("in demo",a)
// }

// func main(){
// 	fmt.Println("in main",a)
// 	demo()
// }




// function scope
// If we want a function to get used outside the package then we need to define that function in Pascal case not in camel case
// In fmt package Println(), Scanln() and toLarge() functions are present 
// 													but we can only use outside fmt Println(), Scanln() function but not toLarge()



// math package (to do mathematical operation using functions present inside it)
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main(){
// 	num := 9.0
// 	res := math.Sqrt(num) // Sqrt() function do operation only on float variable because square root may not always int
// 	fmt.Println("the square root of",num,"is",res)
// }


// If we want to print floating point upto fixed floating point range
// we need to use Printf() function and need to give %.nf here n is total count of floating value
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main(){
// 	num := 10.0
// 	res := math.Sqrt(num)
// 	fmt.Printf("the square root of %.1f is : %.2f", num, res)
// 	fmt.Print("\n")
// }




// conditional statement(if else, switch)(Based on the condition we can change the flow of the program)
// package main

// import "fmt"

// func main(){
// 	num := 8

// 	if num<5{
// 		fmt.Println("Hi", num)
// 	}else {									// else block should start from where if block ends
// 		fmt.Println("Bye", num)
// 	} 
// }


// Print numbers in words(we are checking only for 1 and 2)(if, else if, else)
// package main

// import "fmt"

// func main(){
// 	num :=6

// 	if num == 1 {
// 		fmt.Println("One")
// 	} else if num == 2 {
// 		fmt.Println("Two")
// 	} else{
// 		fmt.Println("None matched")
// 	}
// }




// Assignment (even or odd)
// package main

// import "fmt"

// func main(){
// 	num := 111

// 	if num % 2 == 0 {
// 		fmt.Println("even")
// 	} else {
// 		fmt.Println("odd")
// 	}
// }




// switch (here we dont use break keyword to comes out of switch block, if case matched then execute it and coes out of switch automatically)
// package main

// import "fmt"

// func main(){
// 	num :=9
// 	switch num{
// 	case 1 : fmt.Println("One")
// 	case 2 : fmt.Println("Two")
// 	default : fmt.Println("None matched")
// 	}
// }




// time package (To check when is Saturday)
package main

import (
	"fmt"
	"time"
)

func main(){
	today :=time.Now().Weekday()
	switch time.Saturday {
	case today + 0 :
		fmt.Println("Today is Saturday")
	case today + 1 :
		fmt.Println("Tomorrow is Saturday")
	case today + 2 :
		fmt.Println("Day after tomorrow is Saturday")
	default:
		fmt.Println("Saturday is too far")
	}
}