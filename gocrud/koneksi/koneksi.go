package koneksi

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func DbConn()(db *sql.DB)  {
	dbDriver := "mysql"
	dbUser := "root"
	dbPassword := ""
	dbName := "goblog"

	db, err := sql.Open(dbDriver, dbUser +":"+dbPassword+"@(192.168.64.2)/"+dbName)
	if err != nil{
		fmt.Println(err)
	}
	return db
}