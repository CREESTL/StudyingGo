package side_packages

import (
	"fmt"
	"math"
)

// типы (Circle, Rectangle) можно объвлять ЕДИНОЖДЫ в пакете!!!

// так задается новый тип, эквивалентный структуре
type Circle struct{
	x, y, r float64
}

func createCircle()Circle{
	// здесь все поля инициализируются нулями
	empty := new(Circle)
	_ = empty
	// а так задаются параметры сразу (В ФУГУРНЫХ СКОБКАХ!)
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r)
	return c
}

func circleArea(c Circle)float64{
	return math.Pi * c.r * c.r
}

// ================================================
// некой пародией на ООП является метод-функция
// в ней указывается структура-получатель, к которой данный метод привязывается
// это позволяет вызывать метод через структура.метод
func (c *Circle) area() float64{
	return math.Pi * c.r * c.r
}

func CircleOperations(){
	c := createCircle()
	fmt.Println(circleArea(c))
	fmt.Println(c.area())
}

// то же самое с прямоугольником
type Rectangle struct{
	x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64{
	return (r.x2 - r.x1) * (r.y2 - r.y1)
}

func RectangleOperations(){
	r := Rectangle{0,0,5, 3}
	fmt.Println(r.area())

}


// ================================================
// ВСТРАИВАЕМЫЕ СТРУКТУРЫ
// есть структура человек
type Person struct{
	name string
}
// метод, с помощью которого человек говорит
func (p *Person) talk(){
	fmt.Println("Hello T1000!")
}
// андроид - тоже человек
type Android struct{
	Person
	model string
}

func BeepAndroid(){
	b500 := new(Android)
	// можно обратиться к человеку внутри андроида
	b500.Person.talk()
	// но и сам андроид может говорить!
	b500.talk()
}
