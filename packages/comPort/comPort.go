package comPort

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
)


func New(name string, baud int) <- chan string{
	c := &serial.Config{Name: name, Baud: baud}
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello world")
	ch := make(chan string, 10)

	reader := bufio.NewReader(s)
	go func(){
		for{
			str, err := reader.ReadString('.')
			if err != nil {
				log.Fatal(err)
			}
			ch <- str
		}
	}()
	return ch
}

