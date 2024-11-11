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
			//先监控删除
			errD := DeleteFilesNotInSource(Path1, Path2)
			if errD != nil {
				return
			}

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
	srcFile, err5 := os.Open(FilePath1)
	if err5 != nil {
		fmt.Println(err5)
	}
	defer func(srcFile *os.File) {
		err6 := srcFile.Close()
		if err6 != nil {
			fmt.Println(err6)
		}
	}(srcFile)

	// 创建目标文件
	dstFile, err7 := os.Create(FilePath2)
	if err7 != nil {
		fmt.Println(err7)
	}
	defer func(dstFile *os.File) {
		err8 := dstFile.Close()
		if err8 != nil {

		}
	}(dstFile)

	// 复制文件内容
	_, err9 := io.Copy(dstFile, srcFile)
	if err9 != nil {
		fmt.Println(err9)
	}

	return nil
}

func DeleteFilesNotInSource(path1 string, path2 string) error {

	// 获取path1文件信息
	path1Files, err10 := os.ReadDir(path1)
	if err10 != nil {
		return fmt.Errorf("无法读取源目录: %v", err10)
	}

	// 快速查找(好方法！用csdn找到类似的嘿嘿)
	FilesMap := make(map[string]int)
	for _, file := range path1Files {
		//存在的文件给个1
		FilesMap[file.Name()] = 1
	}

	// 获取path2文件信息
	path2Files, err11 := os.ReadDir(path2)
	if err11 != nil {
		return fmt.Errorf("无法读取目标目录: %v", err11)
	}

	// 遍历目标目录中的文件
	for _, file := range path2Files {

		// 如果目标文件在path1(map)中不存在，则删除目标文件
		if _, exists := FilesMap[file.Name()]; !exists {

			filePath := filepath.Join(path2, file.Name())

			err12 := os.Remove(filePath) // 删除文件
			if err12 != nil {
				fmt.Printf("删除文件 %s 时出错: %v\n", filePath, err12)
			} else {
				fmt.Printf("已删除文件: %s\n", filePath)
			}
		}
	}

	return nil
}
