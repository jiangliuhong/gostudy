package gensql

import (
	"fmt"
	"sqlstat/conn"
	"testing"
)

type Columns struct {
	TableSchema string
	TableName   string
	ColumnName  string
	DataType    string
}

func TestConvertJSON2Text(t *testing.T) {
	db, err := conn.NewMysqlDB("admin:Yunjikeji#123@tcp(cpg-db-dev.yun.idcos:3306)/cloudmsp")
	if err != nil {
		t.Fatal(err)
	}
	sql := `select TABLE_SCHEMA,TABLE_NAME,COLUMN_NAME,DATA_TYPE
from information_schema.COLUMNS 
where  DATA_TYPE = 'json'
and TABLE_SCHEMA in ('ydduc', 'cloudmsp', 'catalog', 'workflow', 'gateway', 'kylin')
order by TABLE_SCHEMA;
`
	rows, err := db.Query(sql)
	if err != nil {
		t.Fatal(err)
	}

	cm := map[string][]Columns{}
	for rows.Next() {
		var tableSchema, tableName, columnName, dataType string
		e := rows.Scan(&tableSchema, &tableName, &columnName, &dataType)
		if e != nil {
			t.Fatal(err)
		}

		var c Columns
		c.TableSchema = tableSchema
		c.TableName = tableName
		c.ColumnName = columnName
		c.DataType = dataType

		//marshal, _ := json.Marshal(c)
		//fmt.Println(string(marshal))

		cm[c.TableSchema] = append(cm[c.TableSchema], c)
	}
	rows.Close()

	// 输出sql
	for k, cs := range cm {
		fmt.Println(k)
		for _, c := range cs {
			text := fmt.Sprintf("ALTER TABLE %s MODIFY COLUMN %s TEXT;", c.TableName, c.ColumnName)
			fmt.Println(text)
		}
	}
}
