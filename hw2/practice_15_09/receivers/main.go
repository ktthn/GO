package main

import "fmt"

func main() {
	rectangle := Rectangle{
		Width:  5,
		Height: 3,
	}

	//a := 5 //просто пример как объявлять пойнтер и ссылать на что-то
	//var pointer *int
	//pointer = &a
	//fmt.Println(a, pointer)
	//fmt.Println(a, *(&a))

	rectangle.Area()
	rectangle.Perimeter()
	rectangle.setWidth(10)
	rectangle.AppendToCalculatedAreas(45)
	fmt.Println(rectangle.CalculatedAreas)

	rectangle.EditFirstArea(90)
	fmt.Println(rectangle.CalculatedAreas)

}

type Rectangle struct {
	Width  int
	Height int

	CalculatedAreas []int
}

func (r *Rectangle) Area() int { //у метода ареа есть прямая ссылка объекта ректжнгл
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() int {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) setWidth(width int) { //у метода ареа есть копия объекта ректэнгл
	r.Width = width
}

func (r *Rectangle) AppendToCalculatedAreas(area int) {
	r.CalculatedAreas = append(r.CalculatedAreas, area)

}

func (r Rectangle) EditFirstArea(area int) { //даже если мы тут передали копию объекта мы смогли поменять площадь
	if len(r.CalculatedAreas) != 0 {
		r.CalculatedAreas[0] = area // слайс под капотом
	}
}
