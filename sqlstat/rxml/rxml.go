package rxml

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type MapperItem struct {
	XMLValue string `xml:",innerxml"`
	Id       string `xml:"id,attr"`
}

type MapperItemSelect struct {
	MapperItem
	XMLName xml.Name `xml:"select"`
}
type MapperItemUpdate struct {
	MapperItem
	XMLName xml.Name `xml:"update"`
}
type MapperItemDelete struct {
	MapperItem
	XMLName xml.Name `xml:"delete"`
}
type MapperItemInsert struct {
	MapperItem
	XMLName xml.Name `xml:"insert"`
}

type Mapper struct {
	XMLName xml.Name           `xml:"mapper"`
	Selects []MapperItemSelect `xml:"select"`
	Updates []MapperItemUpdate `xml:"update"`
	Inserts []MapperItemInsert `xml:"insert"`
	Deletes []MapperItemDelete `xml:"delete"`
}

type SqlStat struct {
	Path       string
	ComplexSql []MapperItem
	SimpleSql  []MapperItem
	Fs         []string
}

func (mapper Mapper) EachSQL(f func(mapperItem MapperItem) error) error {
	for _, item := range mapper.Selects {
		err := f(item.MapperItem)
		if err != nil {
			return err
		}
	}
	for _, item := range mapper.Updates {
		err := f(item.MapperItem)
		if err != nil {
			return err
		}
	}
	for _, item := range mapper.Inserts {
		err := f(item.MapperItem)
		if err != nil {
			return err
		}
	}
	for _, item := range mapper.Deletes {
		err := f(item.MapperItem)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadMapper(path string) (*Mapper, error) {
	xmlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var mapper Mapper
	err = xml.Unmarshal(xmlFile, &mapper)
	if err != nil {
		return nil, err
	}
	return &mapper, nil
}

func CheckComplexSql(sql string) (bool, error) {
	// sql包含高级函数、left join等关系查询关键字
	fs, err := ReadComplexFunc()
	if err != nil {
		return false, err
	}
	for _, name := range fs {
		reg := returnComplexFunReg(name)
		matches := reg.FindStringSubmatch(sql)
		if len(matches) > 0 {
			return true, nil
		}
	}
	return false, nil
	// sql行数大于10行
	//lineCount := 0
	//for i := 0; i < len(sql); i++ {
	//	if sql[i] == '\n' {
	//		lineCount++
	//	}
	//}
	//return lineCount > 40, nil
}

func ReadSqlStat(path string) (ss SqlStat, err error) {
	ss = SqlStat{Path: path}
	mapper, err := ReadMapper(path)
	if err != nil {
		return
	}
	sqlfs := map[string]bool{}
	err = mapper.EachSQL(func(item MapperItem) error {
		sql := item.XMLValue
		fs := resolveSqlF(sql)
		for _, f := range fs {
			sqlfs[f] = true
		}
		complexSql, e := CheckComplexSql(sql)
		if e != nil {
			return e
		}
		if complexSql {
			ss.ComplexSql = append(ss.ComplexSql, item)
		} else {
			ss.SimpleSql = append(ss.SimpleSql, item)
		}
		return nil
	})
	for key, _ := range sqlfs {
		ss.Fs = append(ss.Fs, key)
	}
	return
}
func resolveSqlF(sql string) []string {
	var res []string
	functionRegex := regexp.MustCompile(`\b\w+\s*\(`)
	matches := functionRegex.FindAllString(sql, -1)
	if len(matches) > 0 {
		for _, match := range matches {
			funcName := match
			funcName = strings.TrimSuffix(strings.TrimSpace(funcName), "(")
			funcName = strings.TrimSuffix(strings.TrimSpace(funcName), "\n")
			res = append(res, funcName)

		}
	}
	return res
}

func isReservedWord(word string) bool {
	reservedWords := []string{"values", "from", "select", "where", "and", "or", "insert", "update", "delete"}
	for _, w := range reservedWords {
		if strings.EqualFold(w, word) {
			return true
		}
	}
	return false
}

func ReadComplexFunc() ([]string, error) {
	fi, err := os.Open("./complex_sql.txt")
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	var res []string
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		res = append(res, string(line))
	}
	return res, nil
}

func returnComplexFunReg(name string) *regexp.Regexp {
	str := fmt.Sprintf(`\b%s\s*\(`, name)
	compile := regexp.MustCompile(str)
	return compile
}
