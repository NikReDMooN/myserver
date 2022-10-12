package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func doserver() {
	listener, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		fmt.Println("ошибка", err)
	}
	defer listener.Close()
	fmt.Println("server is working")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(" erros is ", err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	filesdir := "C:\\Myserver\\aaa"
	files, err := ioutil.ReadDir(filesdir)
	if err != nil {
		log.Fatal(err)
	}
	for {
		inputmessage := make([]byte, 1024*4)
		n, err := conn.Read(inputmessage)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}

		handlemeassage := string(inputmessage[0:n])

		switch handlemeassage {
		case "touch":
			fmt.Println("write a file name")
			n, err := conn.Read(inputmessage)
			if n == 0 || err != nil {
				fmt.Println("Read error:", err)
				break
			}
			message := inputmessage[0:n]
			newfilename := string(message)
			os.Create(filesdir + "\\" + newfilename)
			files, _ = ioutil.ReadDir(filesdir)

		case "break":
			fmt.Println("server closed ", handlemeassage)
			break

		case "ls":
			for _, file := range files {
				fmt.Println(file.Name())
			}

		case "cd":
			fmt.Println("write a dictionary")
			n, err := conn.Read(inputmessage)
			if n == 0 || err != nil {
				fmt.Println("Read error:", err)
				break
			}
			newdict := string(inputmessage[0:n])

			filesdir = filesdir + "\\" + newdict
			files, _ = ioutil.ReadDir(filesdir)

		case "mkdir":
			fmt.Println("write a dictionary")
			n, err := conn.Read(inputmessage)
			if n == 0 || err != nil {
				fmt.Println("Read error:", err)
				break
			}
			message := inputmessage[0:n]
			newnamedir := string(message)

			newdir := filesdir + "\\" + newnamedir
			fmt.Println("newd dir = ", newdir)
			_ = os.Mkdir(newdir, 777)
			fmt.Println("dir has been done")
			files, _ = ioutil.ReadDir(filesdir)
		}

		fmt.Println("message is = ", handlemeassage)
	}
}
