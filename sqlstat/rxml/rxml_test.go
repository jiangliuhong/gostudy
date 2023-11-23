package rxml

import (
	"strings"
	"testing"
)

var testMapperPath = "/Users/jiangliuhong/develop/cvs/idcos/cloud-cmp-core/cmp-core-dal/src/main/resources/META-INF/mybatis/sqlmap/CJ_CMP_CIDR_BLOCK.xml"

func TestRead(t *testing.T) {
	path := testMapperPath
	_, err := ReadMapper(path)
	if err != nil {
		t.Logf(err.Error())
		return
	}

}

func TestReadTxt(t *testing.T) {
	sql, err := ReadComplexFunc()
	if err != nil {
		t.Logf(err.Error())
		return
	}
	for _, item := range sql {
		t.Logf(item)
	}
}

func TestCheckComplexSql(t *testing.T) {
	path := testMapperPath
	ss, err := ReadSqlStat(path)
	if err != nil {
		t.Logf(err.Error())
		return
	}
	t.Logf("fs:%s", strings.Join(ss.Fs, ","))
	t.Logf("complexSql:%d;simpleSql:%d", len(ss.ComplexSql), len(ss.SimpleSql))
	//for _, item := range ss.ComplexSql {
	//	t.Logf("id:%s;sql:%s", item.Id, item.XMLValue)
	//}
	t.Log(GenOutput(ss))
}
