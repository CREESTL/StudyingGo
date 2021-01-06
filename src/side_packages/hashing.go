package side_packages

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func SimpleHash(){
	// создание нового хеша
	// IEEE - некриптографическая функция: ее можно обернуть вспять
	h := crc32.NewIEEE()
	// запись строки в хеш
	h.Write([]byte("encrypted_string"))
	// получение контрольной суммы
	v := h.Sum32()
	fmt.Println(v)
}

func getHash(filename string)(uint32, error){
	stringBuffer, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
		return 0, err
	}
	h := crc32.NewIEEE()
	h.Write(stringBuffer)
	return h.Sum32(), nil
}
// если хеш двух строк совпадает, то они с
//большой вероятность одинаковы (но не точно)
//а если отличается - то они точно разные
func CompareStrings(){
	hash1, err := getHash("./src/side_packages/trash/compare1.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	hash2, err := getHash("./src/side_packages/trash/compare2.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(hash1, hash2, hash1 == hash2)
}

// примером криптографической функции может быть алгоритм sha-1
func SHA1(){
	hash := sha1.New()
	hash.Write([]byte("encoded message"))
	// sha1 генерирует хеш в 160 бит, в го нет типа, который мог бы хранить
	// такое число - поэтому используем срез в 20 байт
	buffer := hash.Sum([]byte{})
	fmt.Println(buffer)
}