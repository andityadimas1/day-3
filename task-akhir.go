package main

import "fmt"

func main()  {
	
	// defer func ()  {
	// 	if r :=recover();r!=nil{
	// 		fmt.Println("ran into error")
	// 		main()
	// 	}
	// }()

	functions := map[string]func(int, int)int{
		"+"	: tambah,
		"*"	: kali,
		"-"	: kurang,
		"/" : bagi,		
	}
	var currentNumber int
	fmt.Println("masukkan angka:")
	fmt.Scan(&currentNumber)

	for true{
		var functionName string
		var number int
		fmt.Println("(+ - * /) : ")
		fmt.Scan(&functionName)		
		if functionName == "#"{
			break			
		}
		fmt.Println("Silahkan masukkan angka:")
		fmt.Scan(&number)
		// fmt.Println("Hasilnya:")
		// fmt.Println(currentNumber)
		currentNumber = functions[functionName](currentNumber, number)
	}	
	fmt.Println("Hasilnya:")
		 fmt.Println(currentNumber)
}
func tambah(x,y int) int {
	return x + y
}
func kali(x,y int) int {
	return x * y
}
func kurang(x,y int) int {
	return x - y
}
func bagi(x,y int) int {
	return x / y
}