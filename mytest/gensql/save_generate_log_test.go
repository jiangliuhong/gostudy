package gensql

import (
	"bufio"
	"os"
	"sqlstat/conn"
	"strings"
	"testing"
)

func TestSaveGenerateLog(t *testing.T) {
	db, err := conn.NewMysqlDB("admin:zaq1@WSX@tcp(10.100.96.242:3306)/cloudmsp")
	if err != nil {
		t.Fatal(err)
	}

	sql := `
select argument from mysql.general_log
where event_time between '2023-11-23 15:50:00' and '2023-11-23 18:00:00'
and command_type in ('Query','Prepare','Execute') `
	rows, err := db.Query(sql)
	if err != nil {
		t.Fatal(err)
	}
	var sqls []string
	for rows.Next() {
		var argument string
		e := rows.Scan(&argument)
		if e != nil {
			t.Fatal(err)
		}
		if strings.Index(argument, "select") >= 0 || strings.Index(argument, "SELECT") >= 0 {
			sqls = append(sqls, argument+";")
		}
	}

	file, err := os.Create("query.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	t.Logf("sqls size:%d", len(sqls))
	write := bufio.NewWriter(file)
	for _, content := range sqls {
		write.WriteString(content + "\n")
	}
	write.Flush()
}
