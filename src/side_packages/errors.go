package side_packages

// можно создать новый тип сообщения об ошибке
import (
	"errors"
	"fmt"
	"io/ioutil"
)

// более быстрый способ
func CustomError(){
	_, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(errors.New("This file doesn't exist m8 (("))
		return
	}
}
