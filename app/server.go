package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/myhttp"
)

func handleConnection(conn net.Conn, dir string) {
	defer conn.Close()
	for {

		var buffer []byte = make([]byte, 2048)
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Print("connection closed")
				return
			}
			panic(err)
		}
		message := string(buffer[:n])
		request := myhttp.NewRequest(message)

		fmt.Println(message)
		if request.Line.Url == "/" {
			conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		} else if len(request.Line.Url) > 6 && request.Line.Url[:6] == "/echo/" {
			body := request.Line.Url[6:]

			response := myhttp.NewResponse(body, "text/plain")
			conn.Write([]byte(response.ToString()))

		} else if request.Line.Url == "/user-agent" {
			userAgent := request.Header["User-Agent"]

			response := myhttp.NewResponse(userAgent, "text/plain")
			conn.Write([]byte(response.ToString()))
		} else if len(request.Line.Url) > 7 && request.Line.Url[:7] == "/files/" {
			fileName := request.Line.Url[7:]
			if request.Line.Method == "POST" {
				if len(request.Body) > 0 {
					os.WriteFile(dir+fileName, []byte(request.Body), 0755)
					conn.Write([]byte("HTTP/1.1 201 Created\r\n\r\n"))
					return
				}
			}
			filePath := dir + fileName
			contentBuffer, err := os.ReadFile(filePath)
			if err != nil {
				if os.IsNotExist(err) {
					conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))

				}
				panic(err)
			}
			response := myhttp.NewResponse(string(contentBuffer), "application/octet-stream")
			conn.Write([]byte(response.ToString()))
		} else {
			conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		}

	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	var dir string
	flag.StringVar(&dir, "directory", "", "directory containing the files")
	flag.Parse()
	fmt.Println(dir)

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221", err.Error())
		os.Exit(1)
	}
	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn, dir)
	}
}
