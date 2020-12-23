package math

/*
TODO
Файлы-тесты НЕЛЬЗЯ запускать из main
Они запускаются из КОНСОЛИ командой: go test
TODO
 */

import "testing"


type testpair struct{
	values []float64
	average float64
}

// наборы для тестов
var tests = []testpair{
	{[]float64{}, 0},
	{[]float64{1,2}, 1.5},
	// намеренно указываю не то значение, чтобы тест выдал ошибку
	{[]float64{1,1,1,1}, 10},
	{[]float64{-1,1}, 0},
}

// TODO
// в начале названия функции ОБЯЗАТЕЛЬНО с заглавной буквы Test!
// обязательно в аргументах указывать *testing.T
func TestAverage(t *testing.T){
	for _, pair := range tests{
		v := Average(pair.values)
		if v != pair.average{
			t.Error("For", pair.values, "expected", pair.average, "got", v)
		}
	}
}