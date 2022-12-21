package physics

import "math"

type Parameters struct {
	Height   int // Height in meters
	Angle    int // Measure in degrees
	Velocity int // Velocity in meters/second
	Distance int // meters
}

func CalculateProjectileMotion(parameters Parameters) {
	velocityX := float64(parameters.Velocity) * math.Cos(float64(parameters.Angle))
	velocityY := float64(parameters.Velocity) * math.Sin(float64(parameters.Angle))
}



