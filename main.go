package main

import (
	"fmt"

	"github.com/brunocoronado49/exercies/practices"
)

func main() {
	fmt.Println("Logic Programming")
	fmt.Println("===============================")
	fmt.Println(practices.Vocals("emmanuel"))
	fmt.Println(practices.Vocals("anna"))
	fmt.Println("===============================")
	firstPlayer := []int{2, 7, 4}
	secondPlayer := []int{1, 9, 5}
	fmt.Println(practices.Score(firstPlayer, secondPlayer))
	fmt.Println("===============================")
	fmt.Println(practices.InvertedNumber("54321"))
	fmt.Println("===============================")
	practices.Fizzbuzz(50)
	fmt.Println("===============================")
	fmt.Println(practices.GreetingCityName("Margot Robbie", "Mexico"))
	fmt.Println("===============================")
	fmt.Println(practices.BiggerNumber(5, 4, 7))
	fmt.Println("===============================")
	practices.Weekend()
	fmt.Println("===============================")
	fmt.Println(practices.IsAnagram("listen", "silent"))
	fmt.Println(practices.IsAnagram("dog", "cat"))
	fmt.Println(practices.IsAnagram("mariano", "juan"))
	fmt.Println("===============================")
	practices.Fibonacci(50)
	fmt.Println()
	fmt.Println("===============================")
	fmt.Println(practices.InvertedString("Hello World"))
	fmt.Println("===============================")
}
