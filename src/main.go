package main

/*

Первая строка файла - всегда название пакета!
Есть два типа пакетов:
1) Исполняемые. Это основа приложения, которая и запускается.
В любой программе должен быть пакет main. Го ищет его и начинает выполнение именно с него
2) Дополнительные

--Название не обязательно должно совпадать с папкой (исключение - main)
--В одной папке не могут находится два файла с указанными разными пакетами!!! Если файлы в одной папке - у всех них
должен быть указан один пакет
--Импортировать отдельный ФАЙЛ НЕЛЬЗЯ. Можно только пакет
--Все сущности, которые мы хотим импортировать должны начинаться с БОЛЬШОЙ БУКВЫ
--Наличие неиспользованной переменной или импорта - ОШИБКА (компиляции) - такая философия Go. Он сам удалит
неиспользованный импорт
*/

import (
	"./side_packages"
)
func main(){
	//side_packages.HelloWorld()
	//side_packages.DataTypes()
	//side_packages.CreateVariables()
	//side_packages.InputOutput()
	//side_packages.Cycle()
	//side_packages.EvenOdd()
	//side_packages.NumSwitch()
	//side_packages.ArrayMean()
	//side_packages.Slices()
	//side_packages.Maps()
	//side_packages.CountAverage()
	//side_packages.FunctionsStack()
	//side_packages.SeveralNums()
	//side_packages.CatchError()
	//side_packages.MultiArgs(1, 2, 3, 4, 5)
	//side_packages.FuncInsideFunc()
	//fact := side_packages.Factorial(5)
	//fmt.Println(fact)
	//side_packages.Both()
	//side_packages.ErrorPanic()
	//side_packages.Fib(5)
	//side_packages.CheckZero()
	//side_packages.CheckOne()
	//side_packages.CheckSwap()
	//side_packages.CircleOperations()
	//side_packages.RectangleOperations()
	//side_packages.BeepAndroid()
	//side_packages.InterfaceOperations()
	//side_packages.GoRoutines()
	//side_packages.Channels()
	//side_packages.ShowSelect()
	//side_packages.Sleep(5)
	//side_packages.DRY()
	//side_packages.Strings()
	//side_packages.ReadInStrings()
	//side_packages.ReadSingleSymbol()
	//side_packages.ReadWithScanner()
	//side_packages.ReadFromFile()
	//side_packages.FastReadFromFile()
	//side_packages.WriteIntoFile()
	//side_packages.ListDir()
	//side_packages.RecurrentListDir()
	side_packages.CustomError()
}
