package side_packages

import (
	"container/list"
	"fmt"
)
/*
Помимо списков и карт есть еще некоторые коллекции из пакета container
Здесь представлен только один пример
 */

// двунаправленный список
func LinkedList(){
	var list list.List
	// добавление в список
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	// перебор списка
	for e := list.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}

