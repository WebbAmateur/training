package main

import (
	"flag"

	"family.org/training/vehicle"
)

func main() {
	levelPtr := flag.Int("level", 1, "Level 1..20")
	flag.Parse()

	if *levelPtr < 1 || *levelPtr > 20 {
		panic("Out of range")
	}

	v := vehicle.NewVehicle(*levelPtr)

	vehicle.DumpVehicle(*v)
}
