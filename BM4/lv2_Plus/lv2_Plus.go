package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

//cd ./lv2_Plus
//go run lv2_Plus.go -logfile="example.txt" -loglevel="INFO" -format="txt"

// 确定等级，跟gpt学的好帅的写法嘿嘿（叉腰）
const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

// 全局变量，最后以这个打包
var formattedLog []byte

// 获取级别
func getLogLevelString(level int) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

//在Write方法内打包成这个结构体

type LogFormat struct {
	Timestamp string `json:"timestamp"`
	TimeU     string `json:"timeU"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

// Write方法绑定，最后通过logFile，level和format通过命令行参数确定
type timestampWriter struct {
	logFile io.Writer
	level   int
	format  string
}

func newTimestampWriter(logFile io.Writer, level int, format string) *timestampWriter {
	return &timestampWriter{logFile: logFile, level: level, format: format}
}

// 输出信息(结构体指针方法)
func (t *timestampWriter) Write(p []byte) (err error) {

	TimeNow := time.Now().Format("2006-01-02 15:04:05")

	//搞了好久时间戳www
	TimeUnix := strconv.FormatInt(time.Now().Unix(), 10)

	// 构建日志
	logEntry := LogFormat{
		Timestamp: TimeNow,
		TimeU:     TimeUnix,
		//确定文件等级（命令行参数）
		Level: getLogLevelString(t.level),
		//传入数据为massage
		Message: string(p),
	}

	// 格式输出
	switch strings.ToLower(t.format) {
	case "json":
		logJSON, err := json.Marshal(logEntry)
		if err != nil {
			fmt.Println("Error marshalling log to JSON:", err)
		}
		//json打包
		formattedLog = logJSON
		fmt.Println(string(formattedLog))
	default:
		// txt打包
		formattedLog = []byte(fmt.Sprintf("[%s] [%s] [%s] %s", logEntry.Timestamp, logEntry.TimeU, logEntry.Level, logEntry.Message))
		fmt.Println(string(formattedLog))
	}

	// 写入日志文件
	p = append(formattedLog, '\n')
	if _, err := t.logFile.Write(p); err != nil {
		fmt.Println("Error writing log:", err)
	}
	return nil
}

func main() {
	// 定义命令行参数
	var logFilePath string
	var logLevel string
	var format string
	// 解析命令行参数
	flag.StringVar(&logFilePath, "logfile", "", "文件路径")
	flag.StringVar(&logLevel, "loglevel", "INFO", "日志级别")
	flag.StringVar(&format, "format", "txt", "日志格式(txt和json)")
	flag.Parse()

	//直接获取用户与干了什么
	for {
		formattedLog = nil
		User := ""
		Do := ""
		Exit := ""

		fmt.Println("用户为")
		if _, err1 := fmt.Scanln(&User); err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println("干了什么")
		//为了读取空格，创建缓冲区
		reader2 := bufio.NewReader(os.Stdin)
		Do, _ = reader2.ReadString('\n')

		//加入formattedLog中，在logEntry定义中直接作为Message
		formattedLog = append(formattedLog, User...)
		formattedLog = append(formattedLog, ' ')
		formattedLog = append(formattedLog, Do...)

		// 设置日志级别
		var level int
		switch strings.ToUpper(logLevel) {
		case "DEBUG":
			level = DEBUG
		case "INFO":
			level = INFO
		case "WARN":
			level = WARN
		case "ERROR":
			level = ERROR
		default:
			level = INFO
		}

		// 打开日志文件
		var file io.Writer
		if logFilePath != "" {
			logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			file = logFile
		}

		logWriter := newTimestampWriter(file, level, format)
		if err2 := logWriter.Write(formattedLog); err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println("键入exit退出")
		if _, err3 := fmt.Scanln(&Exit); err3 != nil {
			fmt.Println(err3)
		}
		if Exit == "exit" {
			break
		}
	}

}
