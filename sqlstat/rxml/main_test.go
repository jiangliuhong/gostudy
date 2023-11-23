package rxml

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestComplexSql(t *testing.T) {
	directory := "/Users/jiangliuhong/develop/cvs/idcos"
	var xmlFiles []string
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
	contentMap := map[string]string{}
	complexSqlCount, simpleSqlCount := 0, 0
	//fsMap := map[string]bool{}
	for _, path := range xmlFiles {
		//fmt.Printf("File Path: %s\n", item)
		names := strings.Split(strings.ReplaceAll(path, directory+"/", ""), "/")
		module := names[0]
		if contentMap[module] == "" {
			contentMap[module] = "# 复杂sql统计\n ## " + module + "\n"
		}
		ss, err := ReadSqlStat(path)
		if err != nil {
			t.Fatalf("path:%s;error:%s", path, err.Error())
		}
		complexSqlCount += len(ss.ComplexSql)
		simpleSqlCount += len(ss.SimpleSql)
		//for _, item := range ss.Fs {
		//	fsMap[item] = true
		//}
		output, err := GenOutput(ss)
		if err != nil {
			panic(err)
		}
		contentMap[module] = contentMap[module] + output + "\n"
	}
	for module, content := range contentMap {
		err := flushText(content, "result/"+module+".md")
		if err != nil {
			t.Fatalf("%s:%s", module, err.Error())
		}
	}
	t.Logf("复杂sql:%d;简单sql:%d\n", complexSqlCount, simpleSqlCount)
}

func flushText(content, targetPath string) error {
	csf, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer csf.Close()
	write := bufio.NewWriter(csf)
	write.WriteString(content)
	write.Flush()
	return nil
}

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
		"auto-config.xml",
		"logback-template.xml",
	}
	for _, item := range arr {
		if strings.Index(name, item) >= 0 {
			return true
		}
	}
	return false
}
