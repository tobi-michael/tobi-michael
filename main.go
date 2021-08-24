package main

import "fmt"
type Person struct {
	Age     int
	name    string
	Address string
}

func main() {
	fmt.Println("na me be this")
	strings()
	inter()

	n, a := mich("michael", 500)
	fmt.Println(n, a)
	H := f(15, 12, 4)
	fmt.Println(H)

	FutureAge := calculateAge(15, 13)
	fmt.Println(FutureAge)

	arrays()

	populationlist:= []int{60, 46, 65, 53}
	t:=totalPopulation(populationlist)
	fmt.Println(t )


}

func strings() {
	fmt.Println("my name is Khan")
}

func inter() {
	fmt.Println(47)
}

func mich(name string, population int) (string, int) {
	population = population * 2
	name = "michael"

	return name, population
}

func f(a int, b int, c int) int {

	d := a + b + c
	return d
}

func calculateAge(currentAge, numberOfYears int) int {
	futureAge := currentAge + numberOfYears
	return futureAge
}

func arrays() {
	arra := []int{42, 56, 38, 49}
	for idx, val := range arra {
		fmt.Println(idx, val)
	}

	whyy := []string{"mr", "Niger", "Cad"}
	for idx, val := range whyy {
		fmt.Println(idx, val)
	}

	tobi := Person{
	Address:"lavender",
	Age : 3448983,
	}
	fmt.Println(tobi)
}

func totalPopulation(listOfPopulation []int) int {
	total :=0
	for _, value := range listOfPopulation {
		fmt.Println(total, value)
		total = total + value
	}
	return  total
}