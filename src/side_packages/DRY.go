package side_packages

// DRY - Don't Repeat Yourself
// это принцип повторного использования кода

// импорт указывать до ПАПКИ, в которой лежит файл с нужной нам функцией
import (
	// например эта папка лежит рядом с DRY, а в ней - файл math.go, но его НЕ УКАЗЫВАЕМ в конце
	m "./math"
	"fmt"
)

func DRY() {
	xs := []float64{1,2,3,4}
	avg := m.Average(xs)
	fmt.Println(avg)
}


