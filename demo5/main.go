package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"
)

func main() {
	ch := make(chan string, 100)
	//启动消费者
	for i := 0; i < 3; i++ {
		go func(idx int, ch <-chan string) {
			rand.Seed(time.Now().UnixNano())
			//goIdx := rand.Int()
			for {
				content := <-ch
				fmt.Printf("consumer:%d,reveive:%s", idx, content)
				i := rand.Int63n(1000)
				fmt.Printf("rand number:%d\n", i)
				time.Sleep(time.Duration(i) * time.Millisecond)
			}
		}(i, ch)
	}

	//监听端口
	ln, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(conn)
		for {
			data, err := reader.ReadString('\n')
			if err == io.EOF {
				fmt.Printf("read over:%s", data)
				break
			}

			ch <- data
		}

	}
}
