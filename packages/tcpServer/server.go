package tcpServer

import (
	"bufio"
	"log"
	"net"
	"strings"
)


const (
	addr = "0.0.0.0:5432"
	proto = "tcp4"
)



func New() <- chan string {
	input := make(chan string, 10)
	go func(){
		listener, err := net.Listen(proto, addr)
		if err != nil {
			log.Fatal(err)
		}
		defer close(input)
		defer listener.Close()

		for {
			// Принимаем подключение.
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			// Вызов обработчика подключения.
			handleConn(conn, input)
		}
	}()
	return input

}

func handleConn(conn net.Conn, input chan string) {
	// Закрытие соединения.
	defer conn.Close()
	// Чтение сообщения от клиента.
	reader := bufio.NewReader(conn)
	for{
		b, err := reader.ReadString(';')
		if err != nil {
			log.Println(err)
			return
		}

		// Удаление символов конца строки.
		msg := strings.TrimSuffix(string(b), "\n")
		msg = strings.TrimSuffix(msg, "\r")
		input <- msg
	}
}
