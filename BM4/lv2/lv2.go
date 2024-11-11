package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	var kq []byte
	file, _ := os.OpenFile("lv2/lv2.txt", os.O_WRONLY|os.O_RDONLY, 0666)
	logWriter := newTimestampWriter(file)

	fmt.Fprintln(logWriter.logFile, "登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logWriter.logFile, "用户操作A")
	time.Sleep(3 * time.Second)
	fmt.Fprintln(logWriter.logFile, "操作B")
	logWriter.Write(kq)
}

type timestampWriter struct {
	logFile io.Writer
}

// 把logfile转化成对timestampWriter结构体的指针
func newTimestampWriter(logFile io.Writer) *timestampWriter {
	return &timestampWriter{logFile: logFile}
}

func (t *timestampWriter) Write(p []byte) (n int, err error) {
	TimeNow := []byte(time.Now().Format("2006-01-02 15:04:05"))
	var TimeUnix string
	for i := 0; i < len(TimeNow); i++ {
		p = append(p, TimeNow[i])
	}
	p = append(p, '\n')
	TimeUnix = strconv.FormatInt(time.Now().Unix(), 10)
	p = append(p, TimeUnix...)
	if _, err := t.logFile.Write(p); err != nil {
		fmt.Println(err)
	}
	return len(p), nil
}
