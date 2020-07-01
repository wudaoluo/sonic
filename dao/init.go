package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/wudaoluo/golog"
	"github.com/wudaoluo/sonic/common"
)

var db *xorm.Engine

func Init() {
	storage := common.GetConf().Storage
	var err error
	db, err = xorm.NewEngine("mysql", storage.Addr)
	if err != nil {
		golog.Fatal("mysql连接失败", "err", err)
		panic(err)
	}

	db.ShowSQL(storage.Debug)

	err = db.Ping()
	if err != nil {
		golog.Fatal("mysql ping失败", "err", err)
	}
	golog.Info("mysql连接成功")

}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	golog.Info("DisconnectDB", "func", "关闭mysql数据库")
	if err := db.Close(); nil != err {
		golog.Error("Disconnect from database failed: ", "err", err.Error())
	}
}
