// circle.go in the shapes package
package shapes

import "math"

// Exported function (function name begins by capital letter)
func AreaOfCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

// Unexported function
func diameterOfCircle(radius float64) float64 {
	return 2 * radius
}
