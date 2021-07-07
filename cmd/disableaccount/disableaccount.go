/*
自动禁用到期的账户
*/
package disableaccount

import (
	"fmt"
	"mygoproject/internal/db"
	"mygoproject/internal/getconfig"
	"mygoproject/internal/logger"
	"time"
)

func Disableaccount() {
	today := time.Now().Format("2006-01-02")
	config := getconfig.InitConfigure()
	sql := config.Get("disableaccount.sql").(string)

	var idb db.Database
	idb = new(db.Mysql)
	thisdb := idb.Init()
	ret, err := thisdb.Exec(fmt.Sprintf(sql, today))
	if err != nil {
		logger.Error.Println(err)
	} else {
		n, err := ret.RowsAffected()
		if err != nil {
			logger.Error.Println("get RowsAffected failed,err: ", err)
		} else {
			logger.Info.Println("update success, affected rows: ", n)
		}
	}

	defer thisdb.Close()
}
