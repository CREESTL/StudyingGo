package side_packages

/*

Здесь рассматривается создание переменных

 */

import "fmt"

// имеет глобальную область видимости (через аргумент указывать необязательно)
var out string = "c++"

func CreateVariables(){
	// типа данных указывается ПОСЛЕ имени переменной
	var x string = "Some words"
	fmt.Println(x)
	var s1 string = "First"
	var s2 string = s1 + " and Second"
	var s3 string
	s3 += s2
	fmt.Println(s3)
	//bug int = 32 любая неиспользованная переменная - ошибка (с точки зрения философии Go)

	// АВТОМАТИЧЕСКОЕ ОПРЕДЕЛЕНИЕ ТИПА ПРИ ИНИЦИАЛИЗАЦИИ
	// здесь нет ни var ни типа!
	bug := 32
	// так можно избежать неиспользованной переменной
	_ = bug

	// можно исполдьзовать глобальную переменную
	fmt.Println(out)

	const model string = "T100"

}
// Функция выводит введенное пользователем число
func InputOutput(){
	var input int
	fmt.Println("Enter an integer:")
	// Считывание в переменную
	fmt.Scan(&input)
	fmt.Println(input)

}