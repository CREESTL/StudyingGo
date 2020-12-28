package side_packages

// Здесь рассматривается элемент стандартной библиотеки - строки

import(
	"fmt"
	"strings"
)

func Strings(){
	fmt.Println(
		// true
		strings.Contains("test", "t"),
		// 2
		strings.Count("test", "t"),
		// true
		strings.HasPrefix("test", "te"),
		// false
		strings.HasSuffix("test", "kek"),
		// 1
		strings.Index("test", "e"),
		// a-b
		// добавляет символ-разделитель между строками
		strings.Join([]string{"a", "b"}, "-"),
		// aaaaa
		strings.Repeat("a", 5),
		// bbaaa
		// в заданной строке заменяет букву на новую букву заданное число раз
		strings.Replace("aaaaa", "a", "b", 2),
		// [a, b, c, d, e]
		strings.Split("a-b-c-d-e", "-"),
		// test
		strings.ToLower("TEST"),
		// TEST
		strings.ToUpper("test"),
		)
	// преобразование строки в ASCII (байты) и наоборот
	byteArr := []byte("test")
	fmt.Println(byteArr)
	stringFromBytes := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(stringFromBytes)
}
