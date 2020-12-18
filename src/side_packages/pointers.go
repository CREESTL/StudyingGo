package side_packages

import "fmt"

// функция НЕ меняет значение переменной, потому что внутри создается лишь копия переменной
func zero(x int){
	x = 0
}
// принимает указатель, разыменовывает, обнуляет по адресу
// & возвращает *int
// *int - число
func ptrZero(xPtr *int){
	*xPtr = 0
}
func CheckZero(){
	x := 5
	zero(x)
	// все еще 5
	fmt.Println(x)
	// принимает АДРЕС ячейки (через амперсант)
	fmt.Println(&x)
	ptrZero(&x)
	// вот теперь 0
	fmt.Println(x)

}


func one(xPtr *int){
	*xPtr = 1
}

func CheckOne(){
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

}

// функция меняет числа местами
func swap(xPtr, yPtr *int){
	// в две строки нельзя - не поменяются
	*xPtr, *yPtr = *yPtr, *xPtr
}

func CheckSwap(){
	x, y := 5, 12
	swap(&x, &y)
	fmt.Println(x, y)
}

// в Го есть автоматический сборщик мусора!


