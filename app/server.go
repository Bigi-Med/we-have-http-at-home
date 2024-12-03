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
     strUserAgentValue := strings.Split(strL[2],":")
     path := strLineSplit[1]
     fmt.Printf("HTTP Request Path is : %s \n",path)
     if path == "/" {
         conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
     }else if strings.HasPrefix(path,"/echo"){
         body := strings.Split(path, "/")
         fmt.Printf("BODY MESSAGE IS %s \n",body[2])
         conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-type: text/plain\r\nContent-Length: 3\r\n\r\n"+body[2]))
     } else if path == "/user-agent"{
         fmt.Printf("USER AGENT HEADER %s \n",strUserAgentValue[1])
         conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-type: text/plain\r\nContent-Length: 11\r\n\r\n"+strUserAgentValue[1]))
     }else{
         conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
     }

}
