package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"bytes"
	"net"
)

func main() {
	conn,err := net.Dial("tcp", "192.168.5.20:3000")
	if err != nil {
		fmt.Println("连接错误", err)
	}

	//pack
	msg := "hello i am client"
	re := pack(msg)
	fmt.Println("pack:", string(re))
	conn.Write(re)
	//var data [1024]byte
	//n, err := conn.Read(data[:])
	data := unpack(conn)
	if err != nil {
		fmt.Println("client conn error", err)
	}
	fmt.Println("data:", data)
}

func pack(msg string) []byte {
	fmt.Println("pack")

	msgLen := len(msg)
	length := int16(msgLen)
	//buffer对象
	buf := &bytes.Buffer{}
	//写入包头
	err := binary.Write(buf, binary.BigEndian, length)
	if err != nil {
		fmt.Println("写入包头失败", err)
	}
	data := append(buf.Bytes(), []byte(msg)...)
	return data
}

func unpack(conn net.Conn) string{
	fmt.Println("unpack")
	reader := bufio.NewReader(conn)
	headL,_ := reader.Peek(2)
	buf := bytes.NewBuffer(headL)
	var length int16
	err := binary.Read(buf, binary.BigEndian, &length)
	if err != nil {
		fmt.Println("包头解析失败", err)
		return ""
	}
	if int16(reader.Buffered()) < length+2 {
		fmt.Println("数据解析失败")
		return ""
	}
	data := make([]byte, length+2)
	_,err = reader.Read(data)
	if err != nil {
		fmt.Println("读取数据失败", err)
		return ""
	}
	return string(data[2:])
}
