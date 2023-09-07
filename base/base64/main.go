package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		// 获取第一个输入参数（程序名除外）
		sourcePath := args[1]
		file, err := os.ReadFile(sourcePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		decodedBytes, err := base64.StdEncoding.DecodeString(string(file))
		target := "result.txt"
		targetFile, err := os.OpenFile(target, os.O_RDWR|os.O_CREATE, 0644)
		defer targetFile.Close()
		_, err = targetFile.Write(decodedBytes)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("请提供原文件地址")
	}

}
