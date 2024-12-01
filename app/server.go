package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")


	 l, err := net.Listen("tcp", "0.0.0.0:4221")
	 if err != nil {
	 	fmt.Println("Failed to bind to port 4221")
	 	os.Exit(1)
	 }
	
     conn, err := l.Accept()
	 if err != nil {
	 	fmt.Println("Error accepting connection: ", err.Error())
	 	os.Exit(1)
	 }
     data := make([]byte,1024)
     conn.Read(data)
     str := string(data)
     strL := strings.Split(str, "\r\n")
     strLineSplit := strings.Split(strL[0]," ")
     path := strLineSplit[1]
     fmt.Printf("HTTP Request Path is : %s \n",path)
     if path == "/" {
         conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
     }else{
         conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
     }

}
