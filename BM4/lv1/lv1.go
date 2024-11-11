package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var a []byte
	var b []byte

	file, err := os.Open("lv1/lv1.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	now := time.Now()
	for i := 0; i < 10000; i++ {
		reader.Read(a)
	}
	SinceTime := time.Since(now)
	fmt.Println(SinceTime)

	now = time.Now()
	for i := 0; i < 10000; i++ {
		file.Read(b)
	}
	SinceTime = time.Since(now)
	fmt.Println(SinceTime)

}
