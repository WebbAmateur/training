package main

import (
	"fmt"

	"family.org/training/amort"
)

func main() {

	var principle float64
	principle = 100000.00

	interestRate := 0.1
	const time = 20.0
	const payments = 12.0
	answer := 0.0
	panswer := &answer

	fmt.Printf("Start answer: %f &panswer:%X\n", answer, panswer)

	amort.Calculate(principle, interestRate, time, payments, &answer)

	fmt.Printf("The Amortization of Principal=%.2f  Rate:%.1f  Time:%.0f, Years:=%.0f = $%.2f\n", principle, interestRate, time, time, answer)

	fmt.Printf("End answer: %f &panswer%X\n", answer, panswer)

}

// func Calculate() float64 {

// 	if principle < 0 {
// 		panic(true)
// 	}

// 	numerator := rate * principle
// 	denominator := payments * (1 - math.Pow((1+rate/payments), (-1*payments*time)))
// 	return numerator / denominator
// }
