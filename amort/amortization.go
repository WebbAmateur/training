package amort

import (
	"fmt"
	"math"
)

func Calculate(principle float64, rate float64, time float64, payments float64, pmemory *float64) {
	fmt.Printf("Beginning of Calculate *rho:%f rho:%X\n", *pmemory, pmemory)

	if principle < 0 {
		panic(true)
	}

	numerator := rate * principle
	denominator := payments * (1 - math.Pow((1+rate/payments), (-1*payments*time)))
	*pmemory = numerator / denominator

	fmt.Printf("End of Calculate *rho:%f rho:%X\n", *pmemory, pmemory)
}
