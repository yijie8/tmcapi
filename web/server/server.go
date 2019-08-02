package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type name struct {
	
}
func main() {
	listener,err := net.Listen("tcp","localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for  {
		conn,err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		//有消息就输出出去
		input := bufio.NewScanner(conn)
		for input.Scan() {
			fmt.Println(input.Text())
		}

	}
}
