// package main

// import ("fmt")

// func main(){
// 	fmt.Println("Hello World freeCode camp")
// }

// In Go we can not declare un-used variable(If we declare a variabe we need to use it otherwise it will throw error)
// Variables
// package main

// import ("fmt")

// var packageLevelVar ="this is package level vatiable"
// var packageLevelVar2 string ="this is package level vatiable specifying its type"
// //packageLevelVar3 :="Hi" 			//using := we can not declare package level variables

// // We can declare related variables inside var block so that we do not need to re-write var keyword multiple times
// var(
// 	actorName ="Akshay Kumar"
// 	actorAge =56
// 	actorCareerInYears =20
// 	actorCountry ="India"
// )

// func main(){

// 	var num1 int =10
// 	var num2 = 20
// 	num3 :=40
// 	num3 = 60  // we can change the value or reassign the variables
// 	const num4 =60
// //	num4 = 70		// we can not change or re-assign the value of constants
// 	var num5 uint =80
// //	num5 = -90		// we can not assign or re-assign negative value to unsigned-int type variable
// 	var num6 float32 =10.89 // here we can take float32 or float64
// 	var booleanValue bool = true
// 	var num7 int =100;

// 	fmt.Println(num1, "First way of variable declaration")
// 	fmt.Println()
// 	fmt.Println(num2, "Second way of variable declaration")
// 	fmt.Println()
// 	fmt.Println(num3, "Third way of variable declaration")
// 	fmt.Println()
// 	fmt.Println(num4, "Constant declaration (we can not re-assign to same constant variable)")
// 	fmt.Println()
// 	fmt.Println(num5, "Un-signed-int declaration (we can not declare negative integer here)")
// 	fmt.Println()
// 	fmt.Println(num6, "floating point number (float32, float64)")
// 	fmt.Println()
// 	fmt.Println(booleanValue, ": boolean data type value")
// 	fmt.Println()
// 	fmt.Printf("The value of num7 is %v and its type is %T\n", num7, num7)
// 	fmt.Println()
// 	fmt.Printf("%v and its type is %T\n", packageLevelVar, packageLevelVar)
// 	fmt.Println()
// 	fmt.Printf("%v and its type is %T\n", packageLevelVar2, packageLevelVar2)
// 	fmt.Println()
// 	fmt.Printf("Related Variables from Package level : Name = %v, Age = %v, Acting Career = %v Years, He is From %v",
// 				actorName,actorAge,actorCareerInYears,actorCountry)
// }

// Variable Shadowing(declaring variables with the same name in package level and method level)
// In case of shadowing preference will be given to the method or local variables
// In Go there is no way to declare a private variable however almost same work will do block level variable
// package main

// import ("fmt")

// var num =10;

// func main(){
// 	var num =40
// 	fmt.Println(num, ": because preference will be given to local variable(Shadowing)")
// }

// Variables Scope and Declaration(Block level or Local variable, Package level variable and Global variable)
// for Local and Package level variable we needs to declare its name starting with Lower case
// for Global variable we needs to declare its name starting with Upper case and should be meaningful
// package main

// import ("fmt")

// var numPackageLevel =10;
// var NumGlobalLevel =20;

// func main(){
// 	var numBlockOrLocalLevel =40
// 	fmt.Println(numBlockOrLocalLevel, ": Local level variable name starting with lower case")
// 	fmt.Println()
// 	fmt.Println(numPackageLevel, ": Package level variable name starting with lower case")
// 	fmt.Println()
// 	fmt.Println(NumGlobalLevel, ": Global level variable name starting with Upper case")
// }

// Convert one variable type to another variable type
// int to float
// package main

// import "fmt"

// func main(){
// 	var numI int =20
// 	fmt.Printf("Integer declaration : value %d type is %T \n", numI, numI)
// 	var numf float32
// 	numf = float32(numI) // int to float
// 	fmt.Printf("Converted variable from Integer to float : %f type is %T \n",numf, numf)

// }

// Convert one variable type to another variable type
//float to int (here after conversion we may see some data loss because it will trim to zero after floating point)
// package main

// import "fmt"

// func main(){
// 	var numf float32 =20.65
// 	fmt.Printf("Float declaration : value %f type is %T \n", numf, numf)
// 	var numI int
// 	numI = int(numf) // int to float
// 	fmt.Printf("Converted variable from float to integer : %d type is %T \n",numI, numI)

// }

// Integer to String Conversion
// for this conversion we need to import "strconv" package
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main(){

// 	var numI int = 42
// 	fmt.Printf("Integer declaration : %d type is %T \n",numI, numI)
// 	var intToStr string
// 	intToStr = string(numI) // it will converted to * which ASCII value is 42
// 	fmt.Printf("It will only return the value present at ASCII position 42 and the value is :  %s and type is %T \n", intToStr, intToStr)

// // to convert Integer to String we need to use Itoa() function which is in strconv package
// 	intToStr = strconv.Itoa(numI)
// 	fmt.Printf("After conversion int to string : value %s type is %T \n", intToStr, intToStr)

// }

// Premetive data-types
// 1> boolean type
// 2> Numeric type(integers, floating point, complex numbers)
// 3> text types (string, character)

// boolean type (bool)
// there are two states of bool datatype (true and false)
// default state of bool data type is false

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var bool1  = true
// 	fmt.Printf("value is %v and type is %T\n", bool1, bool1)

// 	var bool2 = false
// 	fmt.Printf("value is %v and type is %T\n", bool2, bool2)
// }

// package main

// import (
// 	"fmt"
// )

// func main(){

// 	var1 := 1==1
// 	fmt.Printf("value is %v and type is %T\n", var1, var1)
// 	var2 := 1==2
// 	fmt.Printf("value is %v and type is %T\n", var2, var2)
// }

// int data type variable
// signed int (int8, int16, int32, int64) if we define int only it will be either int32 or int64 automatically
// un-signed int (uint8, uint16, uint32, uint64) if we define uint only it will be either uint32 or uint64 automatically.
// byte is same as uint8
// cross mathematical operation is not possible, we need to do type conversion explicitly any one of the variable

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var num1 int =12
// 	var num2 int64 = 10
// //  fmt.Println(num1 + num2) // it will throw error
// 	fmt.Println(int64(num1) + num2)
// //											-------------------------or-------------------
// 	fmt.Println(num1 + int(num2))
// }

// Integers operation
// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var num1 =10
// 	var num2 =3
// 	fmt.Println(num1 + num2) // 13
// 	fmt.Println(num1 - num2) // 7
// 	fmt.Println(num1 * num2) // 30
// 	fmt.Println(num1 / num2) // 3
// 	fmt.Println(num1 % num2) // 1
// }

// bit-wise operator
// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var num1 =10 // 1010
// 	var num2 =3  // 0011
// 	fmt.Println(num1 & num2) // 0010 (2)
// 	fmt.Println(num1 | num2) // 1011 (11)
// 	fmt.Println(num1 ^ num2) // 1001 (9)
// 	fmt.Println(num1 &^ num2) // 0100 (8)
// }

// bit-shifting operation
// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var num1 = 8  // 2^3
// 	fmt.Println(num1 << 3) // 2^3 * 2^3 = 2^6 = 64
// 	fmt.Println(num1 >> 3) // 2^3 / 2^3 = 2^0 = 1
// }

// floating point numbers
// we have float32 and float64 in Go
// float by default is float64

// package main

// import "fmt"

// func main(){
// 	floatVal := 3.14
// 	floatVal = 3.7e72  // same as  3.7 * 10^72
// 	floatVal = 3.7E14  //  same as 3.7 * 10^14
// 	fmt.Printf("value is %v and type is %T \n", floatVal, floatVal)
// }

// Arithmetic operations on floating point number
// we dont have remainder operation for float
// we dont have bitwise operation for floating point number
// package main

// import "fmt"

// func main(){
// 	floatVal1 := 6.14
// 	floatVal2 := 3.42
// 	fmt.Printf("Addition : %v", floatVal1+floatVal2)
// 	fmt.Println()
// 	fmt.Printf("Subtraction : %v", floatVal1-floatVal2)
// 	fmt.Println()
// 	fmt.Printf("Multiplication : %v", floatVal1*floatVal2)
// 	fmt.Println()
// 	fmt.Printf("Division : %v", floatVal1/floatVal2)
// }

// Complex number
// there are two types of complex number complex64 and complex128
// default value for complex is complex128
// package main

// import "fmt"

// func main(){
// 	var complexNum = 2 + 10i
// 	fmt.Printf("Value is %v and type is %T ", complexNum, complexNum)
// }

//Addition, Substraction, Multiplication, Division of Compex Number
// package main

// import "fmt"

// func main(){
// 	var cNum1 complex64 = 10 + 2i
// 	var cNum2 complex64 = 5 + 4i

// 	fmt.Println(cNum1 + cNum2)
// 	fmt.Println(cNum1 - cNum2)
// 	fmt.Println(cNum1 * cNum2)
// 	fmt.Println(cNum1 / cNum2)
// }

// // Extracting real and complex numbers
// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var complexNum complex64 = 2 + 10i
// 	fmt.Printf("Value is %v and type is %T \n", real(complexNum), (complexNum))
// 	fmt.Printf("Value is %v and type is %T \n", imag(complexNum), imag(complexNum))
// }

//Complax define another way
// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var complexNum = complex(12, 2)
// 	fmt.Printf("Value is %v and type is %T \n", complexNum, complexNum)
// }

//Text data-type (string and rune)

//String
// Any UTF-8 character in Go is String
// package main

// import "fmt"

// func main(){
// 	var str1 string= "This is String"
// 	fmt.Printf("Value is: %s, Type is: %T", str1, str1)
// }

// we can treat string as an array
// It will return ASCII value of the character.
// package main

// import "fmt"

// func main(){
// 	var str1 string= "This is String"
// 	fmt.Printf("Value is: %v, Type is: %T", str1[5], str1[5])
// }

// If we want to get character we need to do explicit conversion to string
// package main

// import "fmt"

// func main(){
// 	var str1 string= "This is String"
// 	fmt.Printf("Value is: %v, Type is: %T", string(str1[5]), string(str1[5]))
// }

// string concatination
// We can do one Arithmetic operation on string which is concatination
// package main

// import "fmt"

// func main(){
// 	var str1 string= "This is String"
// 	var str2 string= " concatination"
// 	fmt.Printf("Value is: %v, Type is: %T", str1 + str2, str1 + str2)
// }

// we can convert string into collection of bytes or array of bytes
// package main

// import "fmt"

// func main(){
// 	var str1 string= "This is String"
// 	var byteConversion = []byte(str1)
// 	fmt.Printf("Value is: %v, Type is: %T \n", byteConversion, byteConversion)
// }

//rune
// it is similar to char data type in java but in Go rune is int32
// package main

// import "fmt"

// func main(){

// 	var strVal1 string = "a"   // value = a and type is string
// 	fmt.Printf("Value is: %v, Type is: %T \n", strVal1, strVal1)

// 	var strVal2 string = "a"   // value = a and type is string
// 	var strChar = strVal2[0]	// value = 97 and type is uint8
// 	fmt.Printf("Value is: %v, Type is: %T \n", strChar, strChar)

// 	var runeVal1 rune = 'a'   // value = 97 and type is int32
// 	fmt.Printf("Value is: %v, Type is: %T \n", runeVal1, runeVal1)
// }

// Mix-type and multiple variable declaration in one line
// package main

// import "fmt"

// func main(){
// 	var a, b, c, d, e = 10, 10.65, true, "Ajeet", 'a'
// 	fmt.Printf("Value is: %v, Type is: %T \n", a, a)
// 	fmt.Printf("Value is: %v, Type is: %T \n", b, b)
// 	fmt.Printf("Value is: %v, Type is: %T \n", c, c)
// 	fmt.Printf("Value is: %v, Type is: %T \n", d, d)
// 	fmt.Printf("Value is: %v, Type is: %T \n", e, e)
// }

// constants
// Naming of constants : to define constant we need to write const instead of var keyword
//types of the constant is same as type of variables
// we can not change the value of constant further
// package main

// import "fmt"

// func main(){
// 	const constVal int = 10
// 	// constVal = 20	// because of this line it will throw error
// 	fmt.Printf("Value is: %v, Type is: %T \n", constVal, constVal)
// }

// we can declare constant value at compile time only
// we can not declare value to constant type which is evaluated at run time
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main(){
// 	const x = math.Sqrt(4) // it is calculated at run-time, we can not store it in constant type, we can store it in var type only, otherwise it will give error
// 	fmt.Printf("Value is: %v, Type is: %T \n", x, x)

// }

// constant can take all premitive-type values
//we can not declare an array as constant type it will be always var type because arrays are mutable (collections are mutable)
// package main

// import (
// 	"fmt"
// )

// func main(){
// 	const a = 10
// 	const b = 10.34
// 	const c = true
// 	const d = "Ajeet"
// 	fmt.Printf("Value is: %v, Type is: %T \n", a, a)
// 	fmt.Printf("Value is: %v, Type is: %T \n", b, b)
// 	fmt.Printf("Value is: %v, Type is: %T \n", c, c)
// 	fmt.Printf("Value is: %v, Type is: %T \n", d, d)

// }

// constant shadowing
// when we declare a constant with the same name in package and local level
// Preference will be given to the local always
// package main

// import "fmt"

// const a int16 = 10

// func main(){
// 	const a = 20
// 	fmt.Printf("Value is: %v, Type is: %T \n", a, a)
// }

// we can perform mathematical operations between const and var (type should be same) result will be var only
// package main

// import "fmt"

// func main(){
// 	const a int = 20
// 	var b int = 30
// 	fmt.Printf("Value is: %v, Type is: %T \n", a+b, a+b)
// }

// implicit conversion happen for int type in const while performing arithmetic operation with var type result will be var type
// package main

// import "fmt"

// func main(){
// 	const a = 20
// 	var b int16 = 30
// 	fmt.Printf("Value is: %v, Type is: %T \n", a+b, a+b) // result type will be int16 (because of var type) but not for (const a int = 20)
// }

// Enumaration constant iota in Go
// iota value is always 0 and type is int
// we need to initiate iota value with _ (underscore), its write only not read or use
// When we want to increase adjacent constant by 1 then we declare first constant with iota in constant block in package level
// package main

// import "fmt"

// const(
// 	_ = iota
// 	a
// 	b
// 	c
// 	d
// )

// func main(){
// fmt.Printf("value: %v, Type: %T \n", a, a)
// fmt.Printf("value: %v, Type: %T \n", b, b)
// fmt.Printf("value: %v, Type: %T \n", c, c)
// fmt.Printf("value: %v, Type: %T \n", d, d)
// }

// package main

// import "fmt"

// const(
// 	_ = iota	//0
// 	car			//1
// 	truck		//2
// 	jeep		//3
// 	train		//4
// )

// func main(){
// 	var vehicle int = car
// 	fmt.Printf("value: %v, Type: %T \n", vehicle==car, vehicle==car)		// 1==1
// 	fmt.Printf("value: %v, Type: %T \n", vehicle==jeep, vehicle==jeep)		// 1==3
// }

// value of iota increases inside one constant block, if we will take another constant block then value will be re-initialize
// package main

// import "fmt"

// const(
// 	_ = iota	//0
// 	a1			//1
// )

// const(
// 	_ =iota		//0
// 	a2			//1
// )

// func main(){
// 	fmt.Printf("value: %v, Type: %T \n", a1, a1)
// 	fmt.Printf("value: %v, Type: %T \n", a2, a2)
// }

//if we want to skip the increment value then write _ (underscore) there
// package main

// import "fmt"

// const(
// 	_ = iota	//0
// 	a			//1
// 	b			//2
// 	c			//3
// 	_			//4
// 	d			//5
// )

// func main(){
// 	fmt.Println(a, b, c, d)		// 1 2 3 5 (4 will be skipped from here)
// }

// we can perform arithmetic operation on iota it will get applied on all the adjecent constant inside constant block
//  package main

// import "fmt"

// const(
// 	_ = iota + 5	//0 + 5
// 	a			//1 + 5
// 	b			//2 + 5
// 	c			//3	+ 5
// 	d			//4 + 5
// )

// func main(){
// 	fmt.Println(a, b, c, d)
// }

// we can perform bitwise and bit shifting operation
//  package main

// import "fmt"

// const(
// 	_ = iota << 2	//0 << 2
// 	a			//1 << 2
// 	b			//2 << 2
// 	c			//3	<< 2
// 	d			//4 << 2
// )

// func main(){
// 	fmt.Printf("Value: %v, Type: %T \n", a,a)
// 	fmt.Println(a,b,c,d)
// }

// converting memory from Byte to GB
// package main

// import "fmt"

// const(
// 	_ = iota
// 	KB = 1 << (10 * iota)	// 1 * 2^(10*1)
// 	MB						// 2 * 2^(10*2)
// 	GB						// 3 * 2^(10*3)
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

// func main(){
// 	memoryInBytes := 4000000000.0
// 	fmt.Printf("Value after converted to GB is : %.2f GB \n",memoryInBytes/GB )
// 	fmt.Println(KB)
// }

// storing boolean flag inside 1 byte
// package main

// import "fmt"

// const(
// 	isAdmin = 1 << iota
// 	isHeadquarters
// 	canSeeFinancials

// 	canSeeAfrica
// 	canSeeAsia
// 	canSeeEurope
// 	canSeeNorthAmerica
// 	canSeeSouthAmerica
// )

// func main(){
// 	var roles byte = isAdmin | canSeeFinancials | canSeeEurope 	//this way we can store so much boolean flag value in 1 byte(uint8)
// 	fmt.Printf("Value is : %b, Type is : %T \n", roles, roles)

// 	fmt.Printf("Can see Financials? %v \n", canSeeFinancials & roles == canSeeFinancials) // roles has financial access
// 	fmt.Printf("Can see Africa? %v \n", canSeeAfrica & roles == canSeeAfrica) // roles doesnot has Africa access
// }

//Collection(Arrays, Slices)

//Arrays
//Arrays Creation
// arr := [size]type{values saperated by comma}
// arr := [...]int{elements sapereted by comma}
// Arrays can holds only one type of data
// Storing grades of multiple students into a single variable
// Array elements are stored continuously in the memory
// Access of array elements is easy
// package main

// import "fmt"

// func main(){
// 	// grades := []int {58, 34, 68, 100}
// 	grades := [...]int{58, 34, 68, 100} // in this way size will be taken automatically on the basis of number of elements provided
// 	fmt.Printf("Students Grades are : %v", grades)
// }

// empty Array declaration for furthur use, for int value it will initialize with default value which is 0 for string empty
// package main

// import "fmt"

// func main(){
// 	var emptyIntArray [4]int
// 	var emptyStringArray [4]string
// 	fmt.Println(emptyIntArray)  // [0 0 0 0]
// 	fmt.Println(emptyStringArray)  // [   ] -> four spaces
// }

//we can store values in array into particular index
// package main

// import "fmt"

// func main(){
// 	var stringArray [4]string
// 	fmt.Println("Empty string array: ",stringArray)  // [   ] -> four spaces

// 	stringArray[0] = "Ajeet"
// 	stringArray[1] = "Amit"
// 	stringArray[2] = "Suman"
// 	stringArray[3] = "Kumar"
// 	fmt.Println("After Storing values in array: ",stringArray)
// }

// we can get particular indexed value assigned in the array
// package main

// import "fmt"

// func main(){
// 	var stringArray [4]string

// 	stringArray[0] = "Ajeet"
// 	stringArray[1] = "Amit"
// 	stringArray[2] = "Suman"
// 	stringArray[3] = "Kumar"
// 	fmt.Println("Getting value from array present at 2nd index: ",stringArray[2])
// }

// getting size of the arrayusing len(array) function
// package main

// import "fmt"

// func main(){
// 	var stringArray [4]string

// 	fmt.Println("length of the array is: ",len(stringArray))
// }

// Array of Arrays(Multi-dimensional Array)
// package main

// import "fmt"

// func main(){
// 	var identityMatrix [3][3]int = [3][3]int{[3]int{1,2,3}, [3]int{4, 5, 6}, [3]int{7, 8, 9}}
// 	fmt.Println(identityMatrix)
// }

//Array of Arrays(Multi-dimensional Array) .....another way initializing each row individually
// package main

// import "fmt"

// func main(){
// 	var identityMatrix [3][3]int
// 	identityMatrix[0] = [3]int{1,2,3}
// 	identityMatrix[1] = [3]int{4,5,6}
// 	identityMatrix[2] = [3]int{7,8,9}
// 	fmt.Println(identityMatrix)
// }

// Arrays do not pointing same Array in Go if we assign it to any other array, it will create a copy of same array and give to another Array
// package main

// import "fmt"

// func main(){
// 	arr1 := [...]int{1,2,3}
// 	arr2 := arr1		// arr2 is different than arr1, arr2 will have copy of arr1, i will not point to arr1
// 	arr2[1] =5
// 	fmt.Println(arr1) // [1,2,3]
// 	fmt.Println(arr2) // [1,5,3]
// }

// If we want to pointing arr2 to same arr1 then we need to use pointer
// package main

// import "fmt"

// func main(){
// 	arr1 := [...]int{1,2,3}
// 	arr2 := &arr1		// Now arr1 and arr2 will point the same array, now if we will change on either of them it will reflect on both Array
// 	arr2[1] =5
// 	fmt.Println(arr1) // [1,5,3]
// 	fmt.Println(arr2) // [1,5,3]
// }

// Slice
// we do not need to define size of Slice, it will take automatically(it can increase and decrease size based on the element provided to slice)
// package main

// import "fmt"

// func main(){
// 	slice1 := []int{1,2,3}
// 	fmt.Println(slice1)
// 	fmt.Println(len(slice1))
// 	fmt.Println(cap(slice1))
// }

// If we will assign one slice to another slice then both will point to same slice, it will not create copy
// package main

// import "fmt"

// func main(){
// 	slice1 := []int{1,2,3}
// 	slice2 := slice1
// 	slice2[1] =5		// it will reflact on both the slice, since both are representing to same slice
// 	fmt.Println(slice1)
// 	fmt.Println(slice2)
// }

// sub-slicing (we can perform slicing on array as-well)
// package main

// import "fmt"

// func main(){
// 	arr := []int{0,1,2,3,4,5,6,7,8,9,10}
// 	slice1 := arr[:]
// 	slice2 := arr[2:]
// 	slice3 := arr[:8]
// 	slice4 := arr[2:9]

// 	fmt.Println(arr)
// 	fmt.Println(slice1)
// 	fmt.Println(slice2)
// 	fmt.Println(slice3)
// 	fmt.Println(slice4)
// }

// package main

// import "fmt"

// func main(){
// 	arr := []int{0,1,2,3,4,5,6,7,8,9,10}
// 	slice1 := arr[:]
// 	slice2 := arr[2:]
// 	slice3 := arr[:8]
// 	slice4 := arr[2:9]

// 	arr[5] = 100 // it will change the value for all operations

// 	fmt.Println(arr)
// 	fmt.Println(slice1)
// 	fmt.Println(slice2)
// 	fmt.Println(slice3)
// 	fmt.Println(slice4)
// }

//we can use make() function to create slice
// It will take 2 or three arguments
// creating slice with using 2 arguments make(type, size)
// package main

// import "fmt"

// func main(){
// slice1 := make([]int, 4)
// fmt.Println(slice1)
// fmt.Println(len(slice1))
// fmt.Println(cap(slice1))
// }

// creating slice with using 3 arguments make(type, size, capacity)
// we can also define the capacity using 3rd argument which is capacity
// package main

// import "fmt"

// func main(){
// slice1 := make([]int, 4, 100)
// fmt.Println(slice1)
// fmt.Println(len(slice1))
// fmt.Println(cap(slice1))
// }

// we can create a zero capacity slice then we can append elements
// (but it will not append in the same slice,
//	 it will internally create bigger size slice and copy all element to the bigger sized slice)
// if there is lakhs of elements then this can be serious problem to performance of application.
// Better way is to define the capacity at beginning if we know in advance
// package main

// import "fmt"

// func main(){
// slice1 := []int{}					// this line will create slice with zero capacity
// slice1 = append(slice1, 1)	// this line will create another slice with bigger capacity and copy all element from first slice to bigger capacity slice
// fmt.Println(slice1)
// fmt.Println(len(slice1))
// fmt.Println(cap(slice1))
// }

// we can append 1 or more element
// package main

// import "fmt"

// func main(){
// slice1 := []int{}					// this line will create slice with zero capacity
// slice1 = append(slice1, 1,2,3,4)	// first argument we need to pass slice name and then we can pass multiple element to append
// fmt.Println(slice1)
// fmt.Println(len(slice1))
// fmt.Println(cap(slice1))
// }

// we can concatinate two slices to one using spread operator(...)
// package main

// import "fmt"

// func main(){
// slice1 := []int{1,2,3}
// slice2 := []int{4,5,6}
// slice1 = append(slice1, slice2...)	// appending slice2 to slice1
// fmt.Println(slice1)
// fmt.Println(len(slice1))
// fmt.Println(cap(slice1))
// }

//we can remove elements in between of the slice
// package main

// import "fmt"

// func main(){
// slice1 := []int{1,2,3,4,5}
// slice1 = append(slice1[0:2],slice1[3:5]...)	// removing 3 from slice
// fmt.Println(slice1)
// fmt.Println(len(slice1))
// fmt.Println(cap(slice1))
// }

// Maps and Structs

// Maps (It is a key value collection)
// package main

// import "fmt"

// func main(){

// 	cityPopulation := map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 	}

// 	fmt.Println(cityPopulation)
// }

//--------------------------------------or-----------------------------
// package main

// import "fmt"

// func main(){

// 	cityPopulation := make(map[string]int)
// 	cityPopulation = map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 	}

// 	fmt.Println(cityPopulation)
// }

//getting the individual elements from map using its key
// package main

// import "fmt"

// func main(){

// 	cityPopulation := map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 	}

// 	fmt.Println(cityPopulation)
// 	fmt.Println("Population of Mumbai is:",cityPopulation["Mumbai"])
// }

//adding new values in map
// In map order is not preserved
// package main

// import "fmt"

// func main(){

// 	cityPopulation := map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 	}

// 	cityPopulation["Ahmdabad"] = 6003000
// 	fmt.Println(cityPopulation)
// }

//deleting element from map using delete() method
// package main

// import "fmt"

// func main(){

// 	cityPopulation := map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 	}

// 	delete(cityPopulation, "Hyd") // to delete we need to provide map name and any key which we want to delete
// 	fmt.Println(cityPopulation)
// }

// If we are checking for the key which is non present in the map it will return value 0, it may create confusion
// to get rid of this problem we use pop and ok,
// if we get ok value as false and pop value as 0, it means key is not present in the map,
// if we get ok value as true and pop value as 0 then key present with n0 value in the map
// package main

// import "fmt"

// func main(){

// 	cityPopulation := map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 		"Kolkata" : 0,
// 	}

// 	pop1, ok1 := cityPopulation["Patna"]
// 	pop2, ok2 := cityPopulation["Kolkata"]
// 	fmt.Println(cityPopulation)
// 	fmt.Println(pop1, ok1) // it will print 0 false so "Patna" key is not present in the map
// 	fmt.Println(pop2, ok2) // it will print 0 true so "Kolkata" key is present in the map and its population is 0
// }

//if we just want to check if the Entry is present in the map or not then we can use _(Skip variable) at first place
// package main

// import "fmt"

// func main(){

// 	cityPopulation := map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 		"Kolkata" : 0,
// 	}

// 	_, ok := cityPopulation["Kolkata"] // UnderScore is used to skip the value and ok will just check if Kolkata is present or not
// 	fmt.Println(ok)
// }

//We can use len() method to find the numbers of entry present in the map
// package main

// import "fmt"

// func main(){

// 	cityPopulation := map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 		"Kolkata" : 0,
// 	}
// 	fmt.Println(len(cityPopulation))
// }

// Manipulating map in one place will effect in everywhere
// when we store original map in another variable and we manipulate that variable it will reflact manipulation on Original map also
// package main

// import "fmt"

// func main(){

// 	cityPopulation := map[string]int{

// 		"Delhi" : 200000,
// 		"Mumbai" : 800000,
// 		"Hyd" : 600000,
// 		"Blr" : 500000,
// 		"Chennai" : 700000,
// 		"Kolkata" : 0,
// 	}
// 	cityPopulationCopy := cityPopulation // copying original map entries to another variable
// 	delete(cityPopulationCopy, "Hyd")	// deleting "Hyd" from cityPopulationCopy will delete "Hyd" from  cityPopulation also
// 	fmt.Println(cityPopulation)
// }

//--------------------------------------------struct-------------------------------------------------------
// we can use multiple data-types variables in structs
// all other object type stores only one type of data(array and slice, map key is one type only)
// struct gather the related information togather (eg: information of Animal)
// package main

// import "fmt"

// type Animal struct{
// 	name string
// 	canFly bool
// 	canRun bool
// 	numbersOfLegs int
// 	features []string
// }

// func main(){

// 	animal1 := Animal{
// 		name : "Tiger",
// 		canFly : false,
// 		canRun : true,
// 		numbersOfLegs : 4,
// 		features : []string{"Very Fast", "Very Strong", "Eat Meat Only", "Lives in Jungle only"},
// 	}
// 	fmt.Println("Identity Of animal1 : %v \n", animal1) // %v will print values only
// 	fmt.Printf("Identity Of animal1 : %+v \n", animal1) // %+v will print Key as well as values(readability improved)
// }

//we can get one particular value from struct using dot(.)
// package main

// import "fmt"

// type Animal struct{
// 	name string
// 	canFly bool
// 	canRun bool
// 	numbersOfLegs int
// 	features []string
// }

// func main(){

// 	animal1 := Animal{
// 		name : "Tiger",
// 		canFly : false,
// 		canRun : true,
// 		numbersOfLegs : 4,
// 		features : []string{"Very Strong", "Very Fast", "Eat Meat Only", "Lives in Jungle only"},
// 	}
// 	fmt.Printf("name of animal1 is : %v \n", animal1.name)
// 	fmt.Printf("first feature of animal1 is : %v \n", animal1.features[0])
// }

//Another way to define struct is without field-name of struct but we need to provide values in exact same order
// It is not recomended because in future if we change the field name, it may throw error or can retrieve value of another field
// without field name initializer order in main function for struct should be same as we used while creation of struct
// with field name order can be different it does not create any problem
// package main

// import "fmt"

// type Animal struct{
// 	name string
// 	canFly bool
// 	canRun bool
// 	numbersOfLegs int
// 	features []string
// }

// func main(){

// 	animal1 := Animal{
// 		"Tiger",
// 		false,
// 		true,
// 		4,
// 		[]string{"Very Strong", "Very Fast", "Eat Meat Only", "Lives in Jungle only"},
// 	}
// 	fmt.Printf("details of animal1 is : %+v \n", animal1)
// }

// to export struct to another package we needs to declare fields names in Pascal case
// package main

// import "fmt"

// type Animal struct{
// 	Name string
// 	CanFly bool
// 	CanRun bool
// 	NumbersOfLegs int
// 	Features []string
// }

// func main(){

// 	animal1 := Animal{
// 		Name : "Tiger",
// 		CanFly : false,
// 		CanRun : true,
// 		NumbersOfLegs : 4,
// 		Features : []string{"Very Strong", "Very Fast", "Eat Meat Only", "Lives in Jungle only"},
// 	}
// 	fmt.Printf("details of animal1 is : %+v \n", animal1)
// }

//Declaring anonymous structs (very few use)
// we can not use anonymous struct anywhere else
// two pairs of {defining struct}{initializing struct fields}
// package main

// import "fmt"

// func main(){

// 	person1 := struct{name string}{name : "Ajeet"}
// 	fmt.Printf("name of person1 is : %v \n", person1.name)
// }

//unlike map and slices, structs are independent datatype
// if any struct refer to the original struct, changes made in second struct not reflect in original struct
// package main

// import "fmt"

// func main(){

// 	person1 := struct{name string}{name : "Ajeet"}
// 	person2 := person1
// 	person2.name = "Raman"
// 	fmt.Printf("name of person1 is : %v \n", person1.name) // Ajeet
// 	fmt.Printf("name of person2 is : %v \n", person2.name)  // Raman
// }

// if we will use pointer thenn changes in copied structs will reflact in original structs
// package main

// import "fmt"

// func main(){

// 	person1 := struct{name string}{name : "Ajeet"}
// 	person2 := &person1
// 	person2.name = "Raman"
// 	fmt.Printf("name of person1 is : %v \n", person1.name) //Raman
// 	fmt.Printf("name of person2 is : %v \n", person2.name)  // Raman
// }

// Embading
//In Go does not have Inheritance or IS-A relationship it has HAS-A relationship
// package main

// import "fmt"

// type Animal struct{
// 	Name string
// 	Origin string
// }

// type Bird struct{
// 	Animal				// Bird and Animal are totally different things
// 						// (by embading Animal in Bird we can say Bird can have Animal like characteristics)
// 						// embading Animal struct into Bird struct
// 	SpeedKPH float32
// 	CanFly bool
// }

// func main(){

// 	bird1 := Bird{}
// 	bird1.Name ="Emu"
// 	bird1.Origin = "Australia" // in Bird struct we dont have Name and Origin field but due to embading we can use fields in Bird
// 	bird1.SpeedKPH = 48.9
// 	bird1.CanFly = false

// 	fmt.Println(bird1)
// 	fmt.Println(bird1.Name)
// }

//--------------------------------------------------or-------------------------------
// package main

// import "fmt"

// type Animal struct{
// 	Name string
// 	Origin string
// }

// type Bird struct{
// 	Animal
// 	SpeedKPH float32
// 	CanFly bool
// }

// func main(){

// 	bird1 := Bird{
// 		Animal: Animal{Name : "Emu", Origin: "Australia"},
// 		SpeedKPH : 48.9,
// 		CanFly : false,
// 	}

// 	fmt.Println(bird1)
// 	fmt.Println(bird1.Name)
// }

// tag concept in struct(for validation reason eg: name should be maximum of 100 characters) for tag use back tick
// this example is meaningless just for understanding purpose (main use for validation purpose)
// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Animal struct{
// 	Name string `required max: "100"`
// }

// func main(){

// 	t := reflect.TypeOf(Animal{})
// 	field, _ := t.FieldByName("Name")
// 	fmt.Println(field.Tag)

// }
