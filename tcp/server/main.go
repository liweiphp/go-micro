package main

import (
	"fmt"
	"net"
	"bufio"
	"encoding/binary"
	"bytes"
)

func init()  {
	server := &Server{
		ID:"order_1",
		Name: "order",
		Address: "192.168.56.102",
		Port: 3000,
	}
	fmt.Println("注册consul")
	RegisterServer(server)
}

func main() {
	listen,err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		fmt.Println("listen 错误", err)
	}
	fmt.Println("启动tcp服务器 0.0.0.0:3000")
	for  {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("connect error")
		}
		go handler(conn)
	}
}

func handler(conn net.Conn)  {
	defer conn.Close()
	data := unpack(conn)
	fmt.Println("data:", data)
	msg := pack("i am server")
	conn.Write(msg)
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
	//将包长度打包
	err := binary.Read(buf, binary.BigEndian, &length)
	if err != nil {
		fmt.Println("包头解析失败", err)
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