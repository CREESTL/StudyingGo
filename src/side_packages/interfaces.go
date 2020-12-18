package side_packages

import (
	"fmt"
)

// создаются подобно структурам
type Shape interface{
	// внутри перечисляются все методы, которые могут быть реализованы в наследниках
	area() float64
}

// а здесь найдем площадь и круга и прямоугольника
// если бы не было интерфейса Shape, то мы не могли бы указать типа принимаемой переменной
func totalArea(figures ...Shape)float64{
	var area float64
	for _, f := range(figures){
		area += f.area()
	}
	return area
}

// типы (Circle, Rectangle) можно объвлять ЕДИНОЖДЫ в пакете!!!
func InterfaceOperations(){
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 5 ,3}
	area := totalArea(&c, &r)
	fmt.Println(area)
}
