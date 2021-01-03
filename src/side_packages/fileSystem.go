package side_packages

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ПУТЬ ДЛЯ ЧТЕНИЯ НЕОБХОДИМО УКАЗЫВАТЬ ОТНОСИТЕЛЬНО GOPATH (./src/...)

// более "полный способ"
func ReadFromFile(){
	file, err := os.Open("./src/side_packages/trash/test.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	// размер файла
	stat, err := file.Stat()
	if err != nil{
		fmt.Println(err)
		return
	}
	// чтение файла
	byteReader := make([]byte, stat.Size())
	_, err = file.Read(byteReader)
	if err != nil{
		fmt.Println(err)
		return
	}
	str := string(byteReader)
	fmt.Println(str)
}

// более быстрый способ
func FastReadFromFile(){
	bytes, err := ioutil.ReadFile("./src/side_packages/trash/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	strings := string(bytes)
	fmt.Println(strings)
}

// Запись в файл
func WriteIntoFile(){
	file, err := os.Create("./src/side_packages/trash/test_write.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("B.Akunin is the best writer I know!")
}

// Получение имен всех файлов из папки
func ListDir(){
	dir, err := os.Open("./src/side_packages/")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, fileInfo := range fileInfos{
		fmt.Println(fileInfo.Name())
	}
}

// Рекурсивное чтение всех вложенных папок
func RecurrentListDir(){
	// функция, передаваемая в качестве аргумента, вызывается для каждого файла
	filepath.Walk(".", func(path string, info os.FileInfo, err error)error{
		fmt.Println(path)
		return nil
	})
}