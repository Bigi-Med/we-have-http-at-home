package main

import (
	"fmt"
	"net"
	"os"
	"strings"
    "strconv"
)

var _ = net.Listen
var _ = os.Exit

func main() {
    
    args := os.Args[2]

	 l, err := net.Listen("tcp", "0.0.0.0:4221")
     fmt.Println(l)
	 if err != nil {
	 	fmt.Println("Failed to bind to port 4221")
	 	os.Exit(1)
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting connection: ", err.Error())
            os.Exit(1)
        }
        go acceptCon(conn,args)
    }
}
func acceptCon(conn net.Conn, file string){
    data := make([]byte,1024)
    conn.Read(data)
    str := string(data)
    strL := strings.Split(str, "\r\n")
    strLineSplit := strings.Split(strL[0]," ")
    strUserAgentValue := strings.Split(strL[2],":")
    path := strLineSplit[1]
    method := strLineSplit[0]
    fmt.Println(method)
    fmt.Printf("HTTP Request Path is : %s \n",path)
    if method == "POST"{
        fmt.Println(strL[7]) 
        body := strings.Split(path, "/")
        fileName := body[2]
        os.WriteFile(file+fileName,[]byte(strL[7]),0644)
    }else if method == "GET" {

        if path == "/" {
            conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
        }else if strings.HasPrefix(path,"/echo"){
            body := strings.Split(path, "/")
            fmt.Printf("BODY MESSAGE IS %s \n",body[2])
            conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-type: text/plain\r\nContent-Length: 3\r\n\r\n"+body[2]))
        } else if path == "/user-agent"{
            fmt.Printf("USER AGENT HEADER %s \n",strUserAgentValue[1])
            conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-type: text/plain\r\nContent-Length: 11\r\n\r\n"+strUserAgentValue[1]))
        } else if strings.HasPrefix(path,"/files"){
            fmt.Print("Returning file content\n")
            body := strings.Split(path, "/")
            fPath := file+body[2]
            fmt.Println(fPath)
            data, err := os.ReadFile(fPath)
            if err != nil {
                conn.Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\n"))  
            }
            length := len(data)
            conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Length: "+strconv.Itoa(length)+"\r\n\r\n"+string(data)))
        }else{
            conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
        }
     }
 }
