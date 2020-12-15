package side_packages

import "fmt"

// МАССИВЫ
func ArrayMean() {
	// ОБЯЗАТЕЛЬНО УКАЗЫВАТЬ ДЛИНУ!
	// обычная запись массива
	var x [5]int
	x[0] = 5
	x[1] = 3
	x[2] = 3
	x[3] = 2
	x[4] = 5
	// упрощенная запись массива
	y := [5]int{5, 3, 3, 2, 5}
	fmt.Println(x)
	fmt.Println(y)
	var mean int
	// первый способ обхода массива
	for i := 0; i < len(y); i++ {
		mean += y[i]
	}
	mean /= len(x)
	fmt.Println("Mean value1 is", mean)
	mean = 0
	// второй способ обхода массива
	// здесь _ - позиция в массиве (как в enumerate)
	// одиночное подчеркивание - переменная не нужна
	for _, value := range x {
		mean += value
	}
	mean /= len(x)
	fmt.Println("Mean value2 is", mean)
}

// СРЕЗЫ
func Slices(){
	// это те же массивы, только НЕ фиксированной длины
	// пр объявлении длину в квадратных скобках не указываем
	var part[]int32
	_ = part
	// исходный массив
	array := [10]int{1, 16, 75, 26, 4, 657, 2, 1, 96, 666}
	// срез создается функцией make
	// первый аргумент - тип массива, на который срез указывает (к нему он "прикреплен")
	// второй аргумент - длина среза (не больше длины массива)
	// третий аргумент - длина исходного массива
	// чтобы задать пустой срез - второй аргумент = 0
	slice1 := make([]float64, 5, 10)
	_ = slice1
	// а можно записать более коротко
	shortSlice := array[4:8]
	fmt.Println("Short slice is", shortSlice)
	// срез поддерживает дву фунции: append и copy
	//var slice1[]int32
	slice2 := []int32{1,2,3}
	// append применяется вот так!
	slice2 = append(slice2, 4, 5, 6, 7)
	fmt.Println("Slice2 is", slice2)
	// пример работы copy
	slice3 := []int32{1,2,3}
	slice4 := make([]int32, 2)
	// copy применяется вот так!
	copy(slice4, slice3)
	fmt.Println("Slice4 is", slice2)

}


// КАРТЫ
func Maps(){
	// в простонародии - словарь
	// карта: строка-число
	// make - это фактически инициализация
	// длина карты - не огрничена
	var x = make(map[string]int)
	x["ACER"] = 3
	x["HP"] = 5
	x["Lenovo"] = 1
	x["Apple"] = 17
	fmt.Println(x)
	// удаление элемента карты
	delete(x, "ACER")
	fmt.Println(x)
	// проверка удачного извлечения из карты
	if _, ok := x["Apple"]; ok{
		fmt.Println("Value is in the map")
	}
	// можно создавать вложенные карты
	elements := map[string]map[string]string{
		"Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}
	if el, ok := elements["Ne"]; ok{
		if class, ok := el["state"]; ok{
			fmt.Println(el["name"], "is", class)
		}
	}

}
