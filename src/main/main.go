package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	headerSize = 4 // 假设消息头固定为4字节，表示消息长度
)

func main() {
	// 启动服务器
	go startServer()

	// 等待服务器启动
	time.Sleep(time.Second)

	// 启动客户端
	startClient()
}

func startServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	defer listener.Close()
	fmt.Println("Server started, waiting for connections...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection accept failed: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected")

	// 读取消息头
	header := make([]byte, headerSize)
	_, err := conn.Read(header)
	if err != nil {
		log.Printf("Failed to read header: %v", err)
		return
	}

	// 获取消息体长度
	var bodyLength int32
	err = binary.Read(bytes.NewReader(header), binary.BigEndian, &bodyLength)
	if err != nil {
		log.Printf("Failed to read body length: %v", err)
		return
	}

	// 读取消息体
	body := make([]byte, bodyLength)
	_, err = conn.Read(body)
	if err != nil {
		log.Printf("Failed to read body: %v", err)
		return
	}

	fmt.Printf("Received message: %s\n", string(body))
}

func startClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("Client failed to connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("Client connected to server")

	message := "Hello, Server!"
	body := []byte(message)
	bodyLength := int32(len(body))

	// 构造消息头
	header := make([]byte, headerSize)
	err = binary.Write(bytes.NewBuffer(header), binary.BigEndian, bodyLength)
	if err != nil {
		log.Printf("Failed to write header: %v", err)
		return
	}

	// 发送消息
	_, err = conn.Write(header)
	if err != nil {
		log.Printf("Failed to write header: %v", err)
		return
	}
	_, err = conn.Write(body)
	if err != nil {
		log.Printf("Failed to write body: %v", err)
		return
	}

	fmt.Println("Message sent to server", string(header), string(body))
	for {

	}
}
