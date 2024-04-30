// Package clockface provides functions that calculate the positions of the hands.
package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 60
	minutesInHalfClock = 30
	minutesInClock     = 60
	hoursinHalfClock   = 6
	hoursInClock       = 12
)

// A Point represents a two-dimentional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) +
		(math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) +
		(math.Pi / (hoursinHalfClock / float64(t.Hour()%hoursInClock)))
}
