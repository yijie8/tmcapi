package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	var b bytes.Buffer
	b.Write([]byte("您好"))
	_, _ = fmt.Fprint(&b, "http://wwww.baidu.com")
	_, _ = b.WriteTo(os.Stdout)

	return
	conn, err := net.Dial("tcp", "localhost:8000")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	//? 好像是说服务器能联接
	done := make(chan struct{})
	go func() {
		//fmt.Println(1)
		_, _ = io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	///

	ch := make(chan string)
	go clientWriter(conn, ch)
	cin := bufio.NewScanner(os.Stdin)
	for cin.Scan() {
		switch cin.Text() {
		case "register":
			ch <- "register"
		}
	}

	//time.Sleep(3)

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Println(conn, msg)
	}
}
