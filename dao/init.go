package dao

import (
	"bytes"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
)

var db *sql.DB


func InitDB() {
	var buf bytes.Buffer
	var err error

	golog.Info(buf.String())
	db, err = sql.Open("mysql", common.GetConf().Storage.Addr)
	if err != nil {
		golog.Fatal("mysql连接失败", "err",err)
	}

	//设置连接池
	db.SetMaxOpenConns(common.GetConf().Storage.MaxOpen)
	db.SetMaxIdleConns(common.GetConf().Storage.MaxIdle)

	err = db.Ping()
	if err != nil {
		golog.Fatal("mysql ping失败","err", err)
	}
	golog.Info("mysql连接成功")

}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		golog.Error("Disconnect from database failed: " ,"err", err.Error())
	}
}
