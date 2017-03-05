package database

import (
	"github.com/gocraft/dbr"
	_ "github.com/go-sql-driver/mysql"
	
	"api/config/local"//接続する設定ファイルを指定
)

func Init() (*dbr.Session, error) {

	session, err := GetSession()
	
	return session, err
}

func GetSession() (*dbr.Session, error) {
	
	var db_setting string = config.USER + ":" + config.PASSWD + "@tcp(" + config.HOST + ":" + config.PORT + ")/" + config.DB_NAME

	conn, err := dbr.Open("mysql", db_setting, nil)
	
	if err != nil {
		return nil, err
	}else{
		sess := conn.NewSession(nil)
		return sess, nil
	}
}
