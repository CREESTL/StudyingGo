package side_packages

/*
Здесь описывается большинство типов данных, которые есть в Go
 */

/*

--В Go существуют следующие типы целых чисел: uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, NaN
--Строковый тип - string
--Логический тип - true, false (с МАЛЕНЬКОЙ)
--byte (то же самое, что uint8) и rune (то же самое, что int32)

 */

import "fmt"

func DataTypes(){
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println(len("CREESTL"))
	// выводит не сам символ, а ASCII код!
	fmt.Println("CREESTL"[1])
	fmt.Println("Hot" + "Sause")

	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(true || false)
	fmt.Println(!true)


}