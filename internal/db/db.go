package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mygoproject/internal/getconfig"
	"mygoproject/internal/logger"
)

type Database interface {
	Init() *sql.DB
	Query(db *sql.DB, sql string) ([]map[string]string, bool)
}

type Mysql struct {
}

func (f Mysql) Init() *sql.DB {
	config := getconfig.InitConfigure()

	host := config.Get("mysql.host")
	port := config.Get("mysql.port")
	name := config.Get("mysql.name")
	user := config.Get("mysql.user")
	password := config.Get("mysql.password")

	DB, _ := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, password, host, port, name))
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		logger.Error.Println("open database fail")
		panic("open database fail")
	}
	logger.Info.Println("connnect success")
	return DB
}

func (f Mysql) Query(db *sql.DB, sql string) ([]map[string]string, bool) {
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	columns, _ := rows.Columns()
	count := len(columns)
	var values = make([]interface{}, count)
	for i, _ := range values {
		var ii interface{}
		values[i] = &ii
	}
	ret := make([]map[string]string, 0)
	for rows.Next() {
		err := rows.Scan(values...) //这一步就把所有值放到values列表里了
		m := make(map[string]string)
		if err != nil {
			panic(err)
		}
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{}))
			b, _ := raw_value.([]byte)
			v := string(b)
			m[colName] = v
		}
		ret = append(ret, m)
	}

	defer rows.Close()

	if len(ret) != 0 {
		return ret, true
	}
	return nil, false
}
