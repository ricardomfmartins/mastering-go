package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const min = 1
const max = 5

func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type Cube struct {
	x float64
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

type Sphere struct {
	r float64
}

func (c Cube) Vol() float64 {
	return math.Pow(c.x, 3)
}

func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

func (c Sphere) Vol() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.r, 3)
}

type Shapes []Shape3D

func (a Shapes) Len() int {
	return len(a)
}

func (a Shapes) Less(x, y int) bool {
	return a[x].Vol() < a[y].Vol()
}

func (a Shapes) Swap(x, y int) {
	a[x], a[y] = a[y], a[x]
}

func PrintShapes(a Shapes) {
	for _, v := range a {
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown geometric shape.")
		}
	}
	fmt.Println()
}

func main() {
	data := Shapes{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 3; i++ {
		cube := Cube{rF64(min, max)}
		cuboid := Cuboid{rF64(min, max), rF64(min, max), rF64(min, max)}
		sphere := Sphere{rF64(min, max)}
		data = append(data, cube, cuboid, sphere)
	}

	PrintShapes(data)
	sort.Sort(Shapes(data))
	PrintShapes(data)
	sort.Sort(sort.Reverse(Shapes(data)))
	PrintShapes(data)

}
