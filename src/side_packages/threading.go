package side_packages

import (
	"fmt"
	"math/rand"
	"time"
)
//=================================================
// ГОРУТИНЫ
func routine1(n int){
	fmt.Println("f1", n)
	// после каждой операции вывода ожидание от 0 до 250мс (просто чтобы заметить, что все горутины
	// выведутся в консоль одновременно
	nms := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * nms)
}

func GoRoutines(){
	// Горутина - функция, которая пожет выполняться параллельно с отальными
	// чтобы сделать функцию горутиной нужно спереди добавить go
	// программа НЕ будет ждать выполнения всей горутины, а сразу перейдет к следующей строке
	for i:=0; i < 10; i++{
		// горутины очень легкие - их можно создавать кучу
		go routine1(i)
	}
	// сканируем ввод, чтобы программа не завершилась сама
	var stopper string
	fmt.Scanln(&stopper)
}

//=================================================
// КАНАЛЫ

// если прописать (c <- chan string), то из канала можно будет только ДОСТАВАТЬ данные
// а если прописать (c chan <- string), то в канал можно будет только КЛАСТЬ данные
// функция передает в канал одну строку
func ping(c chan string){
	// ТАК можно заменить while True
	for {
		// так делается запись в канал
		c <- "ping"
	}
}

// функция передает в канал другую строку
func pong(c chan string){
	for {
		c <- "pong"
	}
}


func printer(c chan string) {
	for {
		// а так данные вынимаются из канала
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}



func Channels(){
	// Каналы позволяют горутинам общаться друг с другом для синхронизации выполнения операция
	// создается: chan *тип переменных, которые будут передаваться в канале*
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)
	go printer(c)

	var stopper string
	fmt.Scanln(&stopper)

}

/*

SELECT выбирает первый готовый канал. Производит с ним действия. Если готовы два канала, то
первым берет случайный из них. Если ни один канал не готов, то программа приостанавливается
до готовности

 */

func  ShowSelect(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 1)
		}
	}()
	go func(){
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 1)
		}
	}()
	go func(){
		for {
			select{
			case msg1 := <- c1: fmt.Println(msg1)
			case msg2 := <- c2: fmt.Println(msg2)
			}
		}
	}()

	var stopper string
	fmt.Scanln(stopper)
}
