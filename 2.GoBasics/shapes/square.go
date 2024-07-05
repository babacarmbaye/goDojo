// square.go in the shapes package
package shapes

// Exported function
func AreaOfSquare(side float64) float64 {
	return side * side
}

// Unexported function
func perimeterOfSquare(side float64) float64 {
	return 4 * side
}
