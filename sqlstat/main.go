package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func isIgnore(name string) bool {
	arr := []string{
		".idea",
		".git",
		"pom.xml",
		"webapp",
		"mybatis-config.xml",
		"dubbo-config-provider.xml",
		"dubbo-consumer-config.xml",
		"dubbo-config-consumer.xml",
		"spring-consumer.xml",
		"generatorConfig.xml",
		"mybatis-generator.xml",
	}
	for _, item := range arr {
		if strings.Index(name, item) >= 0 {
			return true
		}
	}
	return false
}

func main() {
	directory := "/Users/jiangliuhong/develop/cvs/idcos"
	xmlFiles := []string{}
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if isIgnore(path) {
			return nil
		}
		//fmt.Println(path)
		// 判断是否是文件，并且是 XML 文件
		if !info.IsDir() && filepath.Ext(path) == ".xml" {
			xmlFiles = append(xmlFiles, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("xml file count:%d\n", len(xmlFiles))
	for _, item := range xmlFiles {
		fmt.Printf("File Path: %s\n", item)
	}
}
