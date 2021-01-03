package side_packages

/*
Существует множество пакетов для чтения/вывода в консоль
здесь показан лишь один из них
 */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// чтение строки
func ReadInStrings() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a string")
	for {
		fmt.Print("-> ")
		// указываем до какого символа читать
		text, _ := reader.ReadString('\n')
		// убираем символы конца строки в Windows
		text = strings.Replace(text, "\r\n", "", -1)
		fmt.Println(text)
		if strings.Compare(text, "hi") == 0 {
			fmt.Println("hello, Yourself")
			break
		}
	}
}

// чтение одного символа (руны)
func ReadSingleSymbol(){
	fmt.Println("Input a symbol")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil{
		fmt.Println(err)
	}
	// выводим unicode значение символа
	fmt.Println(char)
}

func ReadWithScanner(){
	fmt.Println("Press 'q' to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		text := scanner.Text()
		fmt.Println(text)
		if text == "q"{
			break
		}

	}
}

