package side_packages

import "fmt"

// В Го всего ОДИН цикл - for!
func Cycle(){
	// первый вариант
	var i int = 1
	// скобки для условия НЕ ставятся
	for i <= 10{
		fmt.Println(i)
		i += 1
	}
	// второй вариант
	for i := 1; i<=10; i++{
		fmt.Println(i)
	}
}

func EvenOdd(){
	for i := 1; i <= 10; i++{
		if i % 2 == 0{
			fmt.Println(i, " is even")
		// else помещается именно так, между двумя скобками
		} else {
			fmt.Println(i, " is odd")
		}
	}
}

func NumSwitch(){
	x := 3
	switch x{
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	}
}
