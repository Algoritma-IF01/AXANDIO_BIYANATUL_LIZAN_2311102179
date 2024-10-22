package main

import("fmt")

func Fibonnaci (n int) int{
	if n == 0{
		return 0
	} else if n == 1{
		return 1
	} else {
		return Fibonnaci(n-1) + Fibonnaci(n-2)
	}
}

func main(){
	fmt.Println("Deret fibonacci hingga suku ke - 10: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Fibonnaci(%d) = %d\n", i , Fibonnaci(i))
	}
}