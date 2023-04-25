package comPort

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
)


func New(name string, baud int) <- chan string{
	c := &serial.Config{Name: name, Baud: baud}
	s, err := serial.OpenPort(c)
	ch := make(chan string, 10)
	if err != nil {
		fmt.Printf("cant open comPort: %v\n", err)
		return ch
	}

	reader := bufio.NewReader(s)
	var str string
	go func(){
		for{
			str, err = reader.ReadString('.')
			if err != nil {
				fmt.Printf("cant read from comPort: %v\n", err)
			}
			ch <- str
		}
	}()
	return ch
}

