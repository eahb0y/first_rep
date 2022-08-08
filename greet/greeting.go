package greet

import (
	"fmt"
)

const pi = 3.14

func PrintCircleArea(radius int) {
	fmt.Printf("Радиус круга %d", radius)
	fmt.Println("Формула падиуса круга 2`2*pi")

	circleArea := CalculateCricleArea(radius)
	fmt.Printf("Плошад круга %f32 = ", circleArea)
}
func CalculateCricleArea(radius int) float32 {
	return float32(radius) * float32(radius) * pi

}
