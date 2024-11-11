package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Message struct {
	Path    string
	ModTime time.Time
	Size    int64
}

//路径一是PRO_MAX_select1,二是PRO_MAX_select2

// select1为被监视路径，select2为自动修改路径
func main() {
	var Path1 string
	var Path2 string
	fmt.Println("请输入路径1")
	if _, err1 := fmt.Scanln(&Path1); err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("请输入路径2")
	if _, err2 := fmt.Scanln(&Path2); err2 != nil {
		fmt.Println(err2)
	}

	MAP1 := GetInFo(Path1)
	MAP2 := GetInFo(Path2)
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("开始监视文件变化...")
			//重新读取信息
			MAP1 = GetInFo(Path1)
			MAP2 = GetInFo(Path2)
			fmt.Printf("当前 MAP1 内容: %+v\n", MAP1)
			fmt.Printf("当前 MAP2 内容: %+v\n", MAP2)
			if !compareMaps(MAP1, MAP2) {
				fmt.Println("文件已发生变化，开始同步...")
				errM := MoveFileInDir(Path1, Path2)
				if errM != nil {
					fmt.Println("文件同步失败:", errM)
					return
				}
				fmt.Println("文件同步完成")
			}

		}
	}
}

//获取文件信息，返回MAP

func GetInFo(Path string) (InFo map[string]Message) {
	var FileName string
	var modTime time.Time
	var size int64
	//用name对应信息
	InFo = make(map[string]Message)
	file, err4 := os.ReadDir(Path)
	if err4 != nil {
		fmt.Println(err4)
	}

	//记录文件信息，加入map
	for _, entry := range file {
		FileName = entry.Name()
		if !entry.IsDir() {
			fileInFo, _ := entry.Info()

			modTime = fileInFo.ModTime()
			size = fileInFo.Size()
			InFo[FileName] = Message{
				Path:    Path,
				ModTime: modTime,
				Size:    size,
			}
		}
	}
	return InFo
}

//人比较笨，只能想出遍历map了QAQ

func compareMaps(map1, map2 map[string]Message) bool {
	if len(map1) != len(map2) {
		return false
	}

	//遍历
	for key, val := range map1 {
		// 检查 map2 中是否有相同的键(搜了搜原来可以这样写嘿嘿)
		if val2, exists := map2[key]; exists {
			// 比较文件的修改时间和大小
			if val.ModTime != val2.ModTime || val.Size != val2.Size {
				return false
			}
		} else {
			// 如果 map2 中没有这个键，说明两个 map 不相等
			return false
		}
	}

	// 如果没有发现任何不匹配的项，则返回 true
	return true
}

//将文件夹下面的全部文件移动

func MoveFileInDir(Path1 string, Path2 string) error {
	File1, _ := os.ReadDir(Path1)
	for _, file1 := range File1 {
		filePath1 := filepath.Join(Path1, file1.Name())
		filePath2 := filepath.Join(Path2, file1.Name())

		if ErrMove := MoveFile(filePath1, filePath2); ErrMove != nil {
			fmt.Println(ErrMove)
		}
	}
	return nil
}

//为MoveFileInDir服务

func MoveFile(FilePath1 string, FilePath2 string) error {
	// 打开源文件
	srcFile, err := os.Open(FilePath1)
	if err != nil {
		return err
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {

		}
	}(srcFile)

	// 创建目标文件
	dstFile, err := os.Create(FilePath2)
	if err != nil {
		return err
	}
	defer func(dstFile *os.File) {
		err := dstFile.Close()
		if err != nil {

		}
	}(dstFile)

	// 复制文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
