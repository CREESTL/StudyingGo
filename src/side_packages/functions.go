package side_packages

import "fmt"

// имя функции, параметры (параметр -> тип), возвращаемый тип (или имя возвращаемого параметра -> тип)
func average(xs []float64) float64{
	total := 0.0
	for _, v := range xs{
		total += v
	}
	return total / float64(len(xs))
}

func CountAverage(){
	var xs = []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	
}

// пример стека вызовов
func FunctionsStack(){
	fmt.Println(f1())
}
func f1()int{
	return f2()
}
// можно явно указать имя возвращаемого параметра
func f2()(result int){
	result = 1
	return
}

// можно возвращать несколько значений
func SeveralNums(){
	x, y := getTwoNums()
	fmt.Println(x ,y)

}

func getTwoNums()(int, int){
	return 5, 6
}

// переменное число аргументов
// аргумент...тип - значит, что количество аргументов такого типа не ограничено
func MultiArgs(args...int){
	sum := 0
	for _, x := range args{
		sum += x
	}
	fmt.Println(sum)
}

// создание функций внутри функций
func FuncInsideFunc(){
	x := 0
	increment := func()int{
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

// пример рекурсии
func Factorial(x int)int{
	if x == 1{
		return x
	}else{
		return x * Factorial(x-1)
	}
}


// ОТЛОЖЕННЫЙ ЗАПУСК
func first(){
	fmt.Println("1st")
}
func second(){
	fmt.Println("2nd")
}
func Both(){
	// функция second выполнится в самом конце функции Both, после всех остальных функций
	defer second()
	first()
}

// ПАНИКА (exception)
// panic просто выводит переданное значение в качетсве ошибки на экран
// recover получает значение, которое было передано в panic и продолжает работу программы
func ErrorPanic(){
	defer func(){
		str := recover()
		fmt.Println(str)
	}() // функция должна вызываться через ()!
	// то есть сначала прога прерывается на панику, потом из-за defer вызывается recover, который просто
	// выводит строку (может делать что угодно)
	panic("ValueError")
}

// упражнение - числа Фибоначчи через рекурсию
func Fib(num int)int{
	if num == 0{
		return 0
	}else if num==1{
		return 1
	}else{
		return Fib(num - 1) + Fib(num - 2)
	}
}