package main

import (
	"net";
	"fmt";
	"bufio"
)


// channel = make(chan string)


func readclient(conn net.Conn) {
  for {
		mensagem, err := bufio.NewReader(conn).ReadString('\n')
   	if err != nil {
			conn.Close()
			break
			// handle error
	  }
		fmt.Print("mensagem: ", string(mensagem) )
    // mandar mensagem para ... memorio compartilha ... , channel
		//conn.Write([]byte("recebido: "+ mensagem))
	}
}

func writeclient(conn net.Conn) {
	for {
	   // ler de mem.comp. ou channel  ;;; bloquear ...
		// msg <- channel_send_client[]
    fmt.Fprintf(conn, "msg")
	}
}

func main(){
	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("escutando na porta 8080!")
	if err != nil {
		// handle error
		// break
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go readclient(conn)
		go writeclient(conn)
	}
}
