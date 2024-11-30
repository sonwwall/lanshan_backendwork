package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// 打开一个日志文件，如果文件不存在则创建，追加写入
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 处理错误
		panic(err)
	}
	defer file.Close()

	// 创建一个带时间戳的写入器
	logWriter := newTimestampWriter(file)

	// 模拟用户操作并记录日志
	fmt.Fprintln(logWriter, "用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作B")

}

// timestampWriter 是一个实现 io.Writer 接口的结构体，它在写入数据前添加时间和时间戳
type timestampWriter struct {
	logFile io.Writer
}

// newTimestampWriter 创建一个 timestampWriter 实例
func newTimestampWriter(w io.Writer) *timestampWriter {
	return &timestampWriter{logFile: w}
}

// Write 方法实现 io.Writer 接口，添加时间戳和时间
func (tw *timestampWriter) Write(p []byte) (n int, err error) {
	// 获取当前时间
	now := time.Now().Format(time.RFC3339)

	// 创建要写入的完整日志信息，包括时间戳
	logMessage := fmt.Sprintf("%s: %s\n", now, string(p))

	// 将日志信息写入到文件
	return tw.logFile.Write([]byte(logMessage))
}

//这个程序有点写不上来，用ai帮了下忙，现在也只是能勉强看懂
