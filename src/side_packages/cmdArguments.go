package side_packages

import (
	"flag"
	"fmt"
	"math/rand"
)

/*

Этот код нигде не вызывается, так как является доп.пакетом и
его неудобно вызывать с командной строки без main

 */
func CMDArgs(){
	// флаг -max, по умолчанию = 6, далее идет описание
	maxp := flag.Int("max", 6, "the max value")
	// после объявления всех флагов их нужно спарсить с командной строки
	flag.Parse()
	// тип флага - *int, поэтому чтобы получить значение - разыменовываем
	fmt.Println(rand.Intn(*maxp))
}