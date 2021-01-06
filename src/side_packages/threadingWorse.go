package side_packages

/*
Обычно параллелизм в го реализуется через горутины и потоки (файл threading.go).
Но есть и более традиционный способ - в пакете sync
 */

// больше информации про мьютексы - здесь: https://ru.stackoverflow.com/questions/661805/%D0%A7%D1%82%D0%BE-%D1%82%D0%B0%D0%BA%D0%BE%D0%B5-%D0%BC%D0%BE%D0%BD%D0%B8%D1%82%D0%BE%D1%80-%D0%BC%D1%8C%D1%8E%D1%82%D0%B5%D0%BA%D1%81-%D0%B8-%D1%81%D0%B5%D0%BC%D0%B0%D1%84%D0%BE%D1%80-%D0%AD%D1%82%D0%BE-%D0%BE%D0%B4%D0%BD%D0%BE-%D0%B8-%D1%82%D0%BE%D0%B6%D0%B5-%D0%B8%D0%BB%D0%B8-%D1%80%D0%B0%D0%B7%D0%BD%D1%8B%D0%B5-%D0%B2%D0%B5%D1%89%D0%B8

import (
	"fmt"
	"sync"
	"time"
)

func Mutex(){
	m := new(sync.Mutex)
	for i := 0; i < 10; i++{
		go func(i int){
			// мьютекст блокирует все потоки, пока не выполнится данный
			// поэтому надписи start-end идут парами, а не врознь
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}
	var input string
	fmt.Scanln(&input)
}