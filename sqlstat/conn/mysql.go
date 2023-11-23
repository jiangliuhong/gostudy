package conn

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// NewMysqlDB
// ds "root:123456@tcp(127.0.0.1:3306)/test"
func NewMysqlDB(ds string) (*sql.DB, error) {
	DB, _ := sql.Open("mysql", ds)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil, err
	}
	return DB, nil
}
