package main

import  (

	"net";
	"fmt";
	"bufio"
)


func main(){

	// conn, err := net.Dial("tcp", "golang.org:80")
	conn, err := net.Dial("tcp", ":80")
	if err != nil {
		// handle error
	}
	for {
		fmt.Fprintf(conn, "GET")
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Fprintf(conn, "teste")

			fmt.Print(status)
		}
	}



}
