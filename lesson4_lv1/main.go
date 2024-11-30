package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	write()
	writeBybufio()

}

func write() {
	fileObj, err := os.OpenFile("test1.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer fileObj.Close()
	str := "泰西小王子\n"
	start := time.Now()
	for i := 0; i < 100000; i++ {
		fileObj.WriteString(str)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("无缓冲写入花费时间：", duration)

}

func writeBybufio() {
	fileObj, err := os.OpenFile("test2.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer fileObj.Close()
	str := "泰西小王子\n"
	writer := bufio.NewWriter(fileObj)
	start := time.Now()
	for i := 0; i < 100000; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("有缓冲写入花费时间：", duration)

}
